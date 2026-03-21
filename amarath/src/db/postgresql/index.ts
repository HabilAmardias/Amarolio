import { Glob, SQL, file } from "bun";

export async function Migrate(db: SQL) {
  const glob = new Glob("./db/migration/*.sql");
  const files = Array.from(glob.scanSync()).sort();
  let query = "";
  for (const filename of files) {
    query += await file(filename).text();
  }
  await db(query).simple();
}

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
