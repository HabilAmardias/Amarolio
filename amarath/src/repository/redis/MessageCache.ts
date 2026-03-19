import { RedisClient } from "bun";
import { Message } from "../../entity/MessageEntity";

export function NewMessageCache(redisClient: RedisClient) {
  return new MessageCacheRepository(redisClient);
}

class MessageCacheRepository {
  redisClient: RedisClient;

  constructor(redisClient: RedisClient) {
    this.redisClient = redisClient;
  }

  SetHistoryContext: (
    chatroomID: string,
    messages: Message[],
  ) => Promise<void> = async (chatroomID, messages) => {
    const EXPIRE_TIME = 60;
    await this.redisClient.set(
      `chatrooms:${chatroomID}:history-context`,
      JSON.stringify(messages),
      "EX",
      EXPIRE_TIME,
    );
  };

  FindHistoryContext: (chatroomID: string) => Promise<Message[] | null> =
    async (chatroomID) => {
      const val = await this.redisClient.get(
        `chatrooms:${chatroomID}:history-context`,
      );
      if (val === null) {
        return val;
      }
      return JSON.parse(val) as Message[];
    };

  RemoveHistoryContext: (chatroomID: string) => Promise<void> = async (
    chatroomID,
  ) => {
    const exists = await this.redisClient.exists(
      `chatrooms:${chatroomID}:history-context`,
    );
    if (exists) {
      await this.redisClient.del(`chatrooms:${chatroomID}:history-context`);
    }
  };
}
