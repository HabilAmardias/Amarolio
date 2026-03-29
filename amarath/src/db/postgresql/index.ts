import { Glob, SQL, file } from "bun";

export async function Migrate(db: SQL) {
  const glob = new Glob("./db/migration/*.sql");
  const files = Array.from(glob.scanSync()).sort();
  for (const filename of files) {
    await db.file(filename);
  }
}

export async function NewPostgresClient() {
  const pgUser = process.env.AMARATH_POSTGRES_USER;
  const pgPassword = process.env.AMARATH_POSTGRES_PASSWORD;
  const pgHost = process.env.AMARATH_DATABASE_HOST;
  const pgDB = process.env.AMARATH_POSTGRES_DB;
  const pgPort = process.env.POSTGRES_PORT;

  const pg = new SQL(
    `postgres://${pgUser}:${pgPassword}@${pgHost}:${pgPort}/${pgDB}`,
  );
  await pg.connect();
  return pg;
}
