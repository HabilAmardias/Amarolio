import Redis from "ioredis";

export async function NewRedisClient(
  password?: string,
  host?: string,
  port?: string,
) {
  const url = `redis://:${password}@${host}:${port}`;
  const rc = new Redis(url);
  await rc.connect();
  return rc;
}
