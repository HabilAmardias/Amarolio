import { RedisClient } from "bun";
import { Chatroom } from "../../entity/ChatroomEntity";

export function NewChatroomCache(redisClient: RedisClient) {
  return new ChatroomCacheRepository(redisClient);
}

class ChatroomCacheRepository {
  redisClient: RedisClient;
  constructor(redisClient: RedisClient) {
    this.redisClient = redisClient;
  }
  FindByID: (chatroomID: string) => Promise<Chatroom | null> = async (
    chatroomID,
  ) => {
    const val = await this.redisClient.get(`chatrooms:${chatroomID}`);
    if (val === null) {
      return val;
    }
    const chatroom = JSON.parse(val) as Chatroom;
    return chatroom;
  };
  SetChatroom: (chatroom: Chatroom) => Promise<void> = async (chatroom) => {
    const EXPIRE_TIME = 60;
    await this.redisClient.set(
      `chatrooms:${chatroom.id}`,
      JSON.stringify(chatroom),
      "EX",
      EXPIRE_TIME,
    );
  };
  DeleteChatroom: (chatroomID: string) => Promise<void> = async (
    chatroomID,
  ) => {
    const exists = await this.redisClient.exists(`chatrooms:${chatroomID}`);
    if (exists) {
      await this.redisClient.del(`chatrooms:${chatroomID}`);
    }
  };
}
