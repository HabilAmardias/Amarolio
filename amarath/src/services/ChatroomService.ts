import { SQL } from "bun";
import {
  Chatroom,
  CreateChatroomParam,
  DeleteChatroomParam,
  GetChatroomsParam,
} from "../entity/ChatroomEntity";
import { NewChatroomRepository } from "../repository/ChatroomRepository";
import { WithTransaction } from "./helper";
import { NewMessageRepository } from "../repository/MessageRepository";
import {
  CustomError,
  ItemNotFoundError,
  UnauthorizedError,
} from "../customerror";

interface OllamaUtilItf {
  CreateTitle: (userMessage: string) => Promise<string>;
}

interface ChatroomCacheItf {
  DeleteChatroom: (chatroomID: string) => Promise<void>;
  SetChatroom: (chatroom: Chatroom) => Promise<void>;
}

interface MessagePublisherItf {
  PublishMessage: (
    channel: string,
    type: string,
    role: string,
    content?: string,
  ) => Promise<number>;
}

interface MessageCacheItf {
  RemoveHistoryContext: (chatroomID: string) => Promise<void>;
}

export function NewChatroomService(
  db: SQL,
  ollamaUtil: OllamaUtilItf,
  messagePublisher: MessagePublisherItf,
  chatroomCache: ChatroomCacheItf,
  messageCache: MessageCacheItf,
) {
  return new ChatroomService(
    db,
    ollamaUtil,
    messagePublisher,
    chatroomCache,
    messageCache,
  );
}

class ChatroomService {
  ollamaUtil: OllamaUtilItf;
  messagePublisher: MessagePublisherItf;
  chatrooomCache: ChatroomCacheItf;
  messageCache: MessageCacheItf;
  db: SQL;
  constructor(
    db: SQL,
    ollamaUtil: OllamaUtilItf,
    messagePublisher: MessagePublisherItf,
    chatroomCache: ChatroomCacheItf,
    messageCache: MessageCacheItf,
  ) {
    this.ollamaUtil = ollamaUtil;
    this.messagePublisher = messagePublisher;
    this.db = db;
    this.chatrooomCache = chatroomCache;
    this.messageCache = messageCache;
  }
  CreateChatroom: (param: CreateChatroomParam) => Promise<string> = async (
    param,
  ) => {
    const ASSISTANT = "assistant";
    const TITLE = "title";

    const chatroomRepo = NewChatroomRepository(this.db);
    const chatroom = await chatroomRepo.CreateChatroom(param.userID, "");
    const channel = `chatrooms:${chatroom.id}:messages`;

    this.ollamaUtil
      .CreateTitle(param.userQuery)
      .catch(() => "New Chat")
      .then(async (val) => {
        chatroom.title = val;
        await Promise.all([
          this.messagePublisher.PublishMessage(channel, TITLE, ASSISTANT, val),
          chatroomRepo.UpdateChatroom(chatroom.id, val),
          this.chatrooomCache.SetChatroom(chatroom),
        ]);
      });

    return chatroom.id;
  };
  GetChatrooms: (param: GetChatroomsParam) => Promise<Chatroom[]> = async (
    param,
  ) => {
    const chatroomRepo = NewChatroomRepository(this.db);
    return await chatroomRepo.GetChatrooms(
      param.userID,
      param.page,
      param.limit,
    );
  };
  DeleteChatroom: (param: DeleteChatroomParam) => Promise<void> = async (
    param,
  ) => {
    await WithTransaction(this.db, async (tx) => {
      const messageRepo = NewMessageRepository(tx);
      const chatroomRepo = NewChatroomRepository(tx);

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

      await messageRepo.DeleteAllMessages(param.chatroomID);
      await chatroomRepo.DeleteChatroom(param.chatroomID);
    });
    await this.messageCache.RemoveHistoryContext(param.chatroomID);
    await this.chatrooomCache.DeleteChatroom(param.chatroomID);
  };
}
