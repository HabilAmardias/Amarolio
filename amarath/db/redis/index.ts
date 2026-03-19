import { RedisClient } from "bun";

export async function NewRedisClient() {
  const url = `redis://:${process.env.REDIS_PASSWORD}@${process.env.REDIS_HOST}:${process.env.REDIS_PORT}`;
  const rc = new RedisClient(url);
  await rc.connect();
  return rc;
}
