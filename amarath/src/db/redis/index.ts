import { RedisClient } from "bun";

export async function NewRedisClient() {
  const url = `redis://:${process.env.AMARATH_REDIS_PASSWORD}@${process.env.AMARATH_REDIS_HOST}:${process.env.REDIS_PORT}`;
  const rc = new RedisClient(url);
  await rc.connect();
  return rc;
}
