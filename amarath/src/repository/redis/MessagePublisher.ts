import { RedisClient } from "bun";

export function NewMessagePublisher(redisClient: RedisClient) {
  return new MessagePublisher(redisClient);
}

class MessagePublisher {
  redisClient: RedisClient;
  constructor(redisClient: RedisClient) {
    this.redisClient = redisClient;
  }
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
}
