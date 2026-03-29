import { RedisClient } from "bun";

export async function NewRedisClient(
  password?: string,
  host?: string,
  port?: string,
) {
  const url = `redis://:${password}@${host}:${port}`;
  const rc = new RedisClient(url);
  await rc.connect();
  return rc;
}
