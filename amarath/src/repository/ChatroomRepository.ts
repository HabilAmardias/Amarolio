import { SQL } from "bun";
import { Chatroom } from "../entity/ChatroomEntity";
import { CustomError, DatabaseExecError } from "../customerror";

export function NewChatroomRepository(dbHandle: SQL) {
  return new ChatroomRepository(dbHandle);
}

class ChatroomRepository {
  dbHandle: SQL;
  constructor(dbHandle: SQL) {
    this.dbHandle = dbHandle;
  }
  CreateChatroom: (userID: string, title: string) => Promise<Chatroom> = async (
    userID,
    title,
  ) => {
    try {
      const [chatroom] = await this.dbHandle<Chatroom[]>`
    INSERT INTO chatrooms (user_id, title)
    VALUES (${userID}, ${title})
    RETURNING id, user_id, title, created_at, updated_at, deleted_at
    `;
      return chatroom;
    } catch (err) {
      throw new CustomError(
        "something went wrong",
        DatabaseExecError,
        (err as Error).message,
      );
    }
  };
  FindByID: (chatroomID: string) => Promise<Chatroom | undefined> = async (
    chatroomID,
  ) => {
    try {
      const [chatroom] = await this.dbHandle<(Chatroom | undefined)[]>`
    SELECT id, user_id, title, created_at, updated_at, deleted_at
    FROM chatrooms
    WHERE id = ${chatroomID} AND deleted_at IS NULL
    `;
      return chatroom;
    } catch (err) {
      throw new CustomError(
        "something went wrong",
        DatabaseExecError,
        (err as Error).message,
      );
    }
  };
  UpdateChatroom = async (chatroomID: string, title?: string) => {
    try {
      let titleUpdate = this.dbHandle`, title = ${title}`;
      await this.dbHandle`
    UPDATE chatrooms
    SET updated_at = NOW() ${title ? titleUpdate : this.dbHandle``}
    WHERE id = ${chatroomID} AND deleted_at IS NULL
    `;
    } catch (err) {
      throw new CustomError(
        "something went wrong",
        DatabaseExecError,
        (err as Error).message,
      );
    }
  };
  GetChatrooms: (
    userID: string,
    page: number,
    limit: number,
  ) => Promise<Chatroom[]> = async (userID, page, limit) => {
    try {
      const chatrooms = await this.dbHandle<Chatroom[]>`
      SELECT id, user_id, title, created_at, updated_at, deleted_at
      FROM chatrooms
      WHERE user_id = ${userID} AND deleted_at IS NULL
      LIMIT ${limit}
      OFFSET ${(page - 1) * limit}
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
  DeleteChatroom: (chatroomID: string) => Promise<void> = async (
    chatroomID,
  ) => {
    try {
      await this.dbHandle`
      UPDATE chatrooms
      SET deleted_at = NOW()
      WHERE id = ${chatroomID} AND deleted_at IS NULL;
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
