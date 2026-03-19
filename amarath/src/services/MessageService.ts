import { NewMessageRepository } from "../repository/MessageRepository";
import { SQL } from "bun";
import { WithTransaction } from "./helper";
import { NewChatroomRepository } from "../repository/ChatroomRepository";
import { GetMessagesParam, SendMessageParam } from "../entity/MessageEntity";
import { Chatroom } from "../entity/ChatroomEntity";
import { Message, MessageResponse } from "../entity/MessageEntity";
import {
  CustomError,
  UnauthorizedError,
  ItemNotFoundError,
  TimeoutError,
} from "../customerror";

interface LLMProviderItf {
  SendMessage: (
    userMessage: string,
    context: string,
    historyMessage?: Message[],
    previousMessage?: string,
    controller?: { signal: AbortSignal },
  ) => Promise<{
    abort: () => void;
    [Symbol.asyncIterator](): AsyncGenerator<
      Awaited<MessageResponse>,
      void,
      unknown
    >;
  }>;
  WebSearch: (userMessage: string) => Promise<string>;
}

interface MessagePublisherItf {
  PublishMessage: (
    channel: string,
    type: string,
    role: string,
    content?: string,
  ) => Promise<number>;
}

interface ChatroomCacheItf {
  FindByID: (chatroomID: string) => Promise<Chatroom | null>;
  SetChatroom: (chatroom: Chatroom) => Promise<void>;
}

interface MessageCacheItf {
  SetHistoryContext: (chatroomID: string, messages: Message[]) => Promise<void>;
  FindHistoryContext: (chatroomID: string) => Promise<Message[] | null>;
}

export function NewMessageService(
  messagePublisher: MessagePublisherItf,
  ollamaUtil: LLMProviderItf,
  chatroomCache: ChatroomCacheItf,
  messageCache: MessageCacheItf,
  db: SQL,
) {
  return new MessageService(
    messagePublisher,
    ollamaUtil,
    chatroomCache,
    messageCache,
    db,
  );
}

class MessageService {
  messagePublisher: MessagePublisherItf;
  ollamaUtil: LLMProviderItf;
  chatroomCache: ChatroomCacheItf;
  messageCache: MessageCacheItf;
  db: SQL;

  constructor(
    messagePublisher: MessagePublisherItf,
    ollamaUtil: LLMProviderItf,
    chatroomCache: ChatroomCacheItf,
    messageCache: MessageCacheItf,
    db: SQL,
  ) {
    this.db = db;
    this.messagePublisher = messagePublisher;
    this.ollamaUtil = ollamaUtil;
    this.chatroomCache = chatroomCache;
    this.messageCache = messageCache;
  }

  SendMessage = async (param: SendMessageParam) => {
    const controller = new AbortController();
    const timeoutID = setTimeout(() => controller.abort(), 120_000);

    try {
      const channel = `chatrooms:${param.id}:messages`;
      const MESSAGE = "message";
      const ASSISTANT = "assistant";
      const CONTINUE = "continue";
      const USER = "user";
      const DONE = "done";

      const chatroomRepo = NewChatroomRepository(this.db);
      const messageRepo = NewMessageRepository(this.db);

      let fullResponse = "";
      let hitLimit = false;

      let chatroom = await this.chatroomCache.FindByID(param.id);
      if (!chatroom) {
        chatroom = await chatroomRepo.FindByID(param.id);
        if (!chatroom) {
          throw new CustomError(
            "item not found",
            ItemNotFoundError,
            "chatroom does not exist",
          );
        }
      }
      if (chatroom.user_id !== param.userID) {
        throw new CustomError(
          "unauthorized access",
          UnauthorizedError,
          `chatroom with id ${chatroom.id} does not belong to the user`,
        );
      }

      let history = await this.messageCache.FindHistoryContext(param.id);
      if (!history) {
        history = await messageRepo.GetMessages(param.id, 5); // get last 5 messages
      }

      const [_, context] = await Promise.all([
        this.messagePublisher.PublishMessage(
          channel,
          MESSAGE,
          USER,
          param.userMessage,
        ),
        this.ollamaUtil.WebSearch(param.userMessage),
      ]);

      const response = await this.StreamMessage(
        channel,
        param.userMessage,
        context,
        history,
        undefined,
        controller,
      );
      fullResponse += response.streamResponse;

      if (response.wasTruncated) {
        await this.messagePublisher.PublishMessage(
          channel,
          CONTINUE,
          ASSISTANT,
        );
        const contResponse = await this.StreamMessage(
          channel,
          param.userMessage,
          context,
          history,
          fullResponse,
          controller,
        );
        fullResponse += contResponse.streamResponse;
        hitLimit = contResponse.wasTruncated;
      }

      await this.messagePublisher.PublishMessage(
        channel,
        DONE,
        ASSISTANT,
        hitLimit ? "TOKEN LIMIT" : "",
      );

      await WithTransaction(this.db, async (tx) => {
        const messageRepo = NewMessageRepository(tx);
        const chatroomRepo = NewChatroomRepository(tx);

        await messageRepo.SaveMessage(param.userMessage, param.id, USER);
        await messageRepo.SaveMessage(fullResponse, param.id, ASSISTANT);
        await chatroomRepo.UpdateChatroom(param.id);
      });

      chatroomRepo.FindByID(param.id).then((val) => {
        if (val) {
          this.chatroomCache.SetChatroom(val);
        }
      });

      messageRepo.GetMessages(param.id, 5).then((val) => {
        this.messageCache.SetHistoryContext(param.id, val);
      });

      return fullResponse;
    } finally {
      clearTimeout(timeoutID);
    }
  };
  GetMessages = async (param: GetMessagesParam) => {
    const chatroomRepo = NewChatroomRepository(this.db);
    const chatroom = await chatroomRepo.FindByID(param.chatroomID);
    if (!chatroom) {
      throw new CustomError(
        "item not found",
        ItemNotFoundError,
        "chatroom does not exist",
      );
    }

    if (chatroom.user_id !== param.userID) {
      throw new CustomError(
        "unauthorized access",
        UnauthorizedError,
        `chatroom with id ${chatroom.id} does not belong to the user`,
      );
    }

    const messageRepo = NewMessageRepository(this.db);
    const messages = await messageRepo.GetMessages(
      param.chatroomID,
      param.limit,
      param.lastID,
    );

    return messages;
  };
  private StreamMessage: (
    channel: string,
    userMessage: string,
    context: string,
    history?: Message[],
    prevMessage?: string,
    controller?: { signal: AbortSignal },
  ) => Promise<{ wasTruncated: boolean; streamResponse: string }> = async (
    channel,
    userMessage,
    context,
    history,
    prevMessage,
    controller,
  ) => {
    let buffer = "";
    let wasTruncated = false;
    let streamResponse = "";

    const MAX_BUFFER_LENGTH = 150;
    const MESSAGE = "message";
    const ASSISTANT = "assistant";
    let inThinking = false;

    const stream = await this.ollamaUtil.SendMessage(
      userMessage,
      context,
      history,
      prevMessage,
      controller,
    );

    for await (const chunk of stream) {
      if (controller?.signal.aborted) {
        stream.abort();
        throw new CustomError(
          "request timeout",
          TimeoutError,
          "request timeout, operation aborted",
        );
      }
      if (chunk.message.thinking && !inThinking) {
        inThinking = true;
        await this.messagePublisher.PublishMessage(
          channel,
          MESSAGE,
          ASSISTANT,
          "Thinking....",
        );
      }
      if (!chunk.message.thinking && chunk.message.content) {
        buffer += chunk.message.content;
        streamResponse += chunk.message.content;
        if (buffer.length > MAX_BUFFER_LENGTH) {
          await this.messagePublisher.PublishMessage(
            channel,
            MESSAGE,
            ASSISTANT,
            buffer,
          );
          buffer = "";
        }
      }
      if (chunk.done_reason === "length") {
        wasTruncated = true;
      }
    }
    if (buffer.length > 0) {
      await this.messagePublisher.PublishMessage(
        channel,
        MESSAGE,
        ASSISTANT,
        buffer,
      );
    }
    return { wasTruncated, streamResponse };
  };
}
