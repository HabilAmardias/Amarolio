import { SQL, TransactionSQL } from "bun";

export function WithTransaction(
  db: SQL,
  callback: (tx: TransactionSQL) => Promise<void>,
) {
  return db.begin<Promise<void>>(callback);
}
