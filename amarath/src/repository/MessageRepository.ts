import { SQL } from "bun";
import { Message, SaveMessageQuery } from "../entity/MessageEntity";
import { CustomError, DatabaseExecError } from "../customerror";

export function NewMessageRepository(dbHandle: SQL) {
  return new MessageRepository(dbHandle);
}

class MessageRepository {
  dbHandle: SQL;
  constructor(dbHandle: SQL) {
    this.dbHandle = dbHandle;
  }
  SaveMessages = async (messages: SaveMessageQuery[], chatroomID: string) => {
    try {
      if (messages.length === 0) {
        return null;
      }

      let query = "INSERT INTO messages(chatroom_id, role, content) VALUES ";
      messages.forEach((val, i) => {
        query += `(${chatroomID}, ${val.role}, ${val.content})`;
        if (i !== messages.length - 1) {
          query += ",";
        }
        query += " ";
      });
      query +=
        "RETURNING id, chatroom_id, role, content, created_at, updated_at, deleted_at";

      return await this.dbHandle<Message[]>(query);
    } catch (err) {
      throw new CustomError(
        "something went wrong",
        DatabaseExecError,
        (err as Error).message,
      );
    }
  };
  GetMessages: (
    chatroomID: string,
    limit: number,
    lastID?: number,
  ) => Promise<Message[]> = async (chatroomID, limit, lastID) => {
    try {
      const MAX_INT = 2_147_483_647;
      const chatrooms = await this.dbHandle<Message[]>`
    WITH chats AS(
      SELECT
        id,
        chatroom_id,
        role,
        content,
        created_at,
        updated_at,
        deleted_at
      FROM messages
      WHERE chatroom_id = ${chatroomID}
      AND id < COALESCE(${lastID}, ${MAX_INT})
      AND deleted_at IS NULL
      ORDER BY created_at DESC
      LIMIT ${limit}
    )
    SELECT *
    FROM chats
    ORDER BY chats.created_at ASC
    `;
      return chatrooms;
    } catch (err) {
      throw new CustomError(
        "something went wrong",
        DatabaseExecError,
        (err as Error).message,
      );
    }
  };
  DeleteAllMessages: (chatroomID: string) => Promise<void> = async (
    chatroomID,
  ) => {
    try {
      await this.dbHandle`
    UPDATE messages
    SET deleted_at = NOW()
    WHERE chatroom_id = ${chatroomID} AND deleted_at IS NULL
    `;
    } catch (err) {
      throw new CustomError(
        "something went wrong",
        DatabaseExecError,
        (err as Error).message,
      );
    }
  };
}
