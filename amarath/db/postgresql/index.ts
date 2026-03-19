import { SQL } from "bun";

export async function NewPostgresClient() {
  const pgUser = process.env.POSTGRES_USER;
  const pgPassword = process.env.POSTGRES_PASSWORD;
  const pgHost = process.env.POSTGRES_HOST;
  const pgDB = process.env.POSTGRES_DB;
  const pgPort = process.env.POSTGRES_PORT;

  const pg = new SQL(
    `postgres://${pgUser}:${pgPassword}@${pgHost}:${pgPort}/${pgDB}`,
  );
  await pg.connect();
  return pg;
}
