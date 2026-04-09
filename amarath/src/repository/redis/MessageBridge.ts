import { randomUUIDv7 } from "bun";
import Redis from "ioredis";

export function NewMessageBridge(redisClient: Redis) {
  return new MessageBridge(redisClient);
}

class MessageBridge {
  redisClient: Redis;
  constructor(redisClient: Redis) {
    this.redisClient = redisClient;
  }
  EnsureGroup = async (stream: string, group: string) => {
    await this.redisClient.xgroup("CREATE", stream, group, "$", "MKSTREAM");
  };
  PublishMessage = async (
    channel: string,
    type: string,
    role: string,
    content?: string,
  ) => {
    return await this.redisClient.publish(
      channel,
      JSON.stringify({
        type,
        role,
        content,
      }),
    );
  };
  AcknowledgeMessage = async (id: string, stream: string, group: string) => {
    await this.redisClient.xack(stream, group, id);
  };
  ReceiveMessage = async (
    group: string,
    stream: string,
    block: string = "5000",
    count: string = "4",
  ) => {
    const id = randomUUIDv7();
    const res = (await this.redisClient.xreadgroup(
      "GROUP",
      group,
      id,
      "COUNT",
      count,
      "BLOCK",
      block,
      "STREAMS",
      stream,
      ">",
    )) as [string, [string, string[]][]][] | null;
    return res;
  };
}
