import { Chatroom } from "../../entity/ChatroomEntity";
import Redis from "ioredis";

export function NewChatroomCache(redisClient: Redis) {
  return new ChatroomCacheRepository(redisClient);
}

class ChatroomCacheRepository {
  redisClient: Redis;
  constructor(redisClient: Redis) {
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
