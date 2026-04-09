import { GetMessagesParam, SendMessageParam } from "../entity/MessageEntity";
import { Message } from "../entity/MessageEntity";
import { ServerResponse } from "../dto";
import { GetMessagesRes } from "../dto/MessageDto";
import Elysia, { t } from "elysia";
import { CustomError, UnauthorizedError } from "../customerror";

interface MessageServiceItf {
  GetMessages: (param: GetMessagesParam) => Promise<Message[]>;
}

export function NewMessageController(messageService: MessageServiceItf) {
  return new Elysia().group("/api/v1", (app) => {
    return app
      .guard({
        headers: t.Object({
          "x-user-id": t.Optional(t.String()),
        }),
      })
      .resolve(({ headers }) => {
        if (!headers["x-user-id"]) {
          throw new CustomError(
            "unauthorized access",
            UnauthorizedError,
            "no user id found",
          );
        }
        return { userID: headers["x-user-id"] };
      })
      .get(
        "/chatrooms/:id/messages",
        async ({ userID, params, query, status }) => {
          const param = {
            userID,
            chatroomID: params.id,
            limit: query.limit,
            lastID: query.last_id,
          };
          const messages = await messageService.GetMessages(param);
          const res: ServerResponse<GetMessagesRes> = {
            success: true,
            data: {
              messages,
            },
          };
          return status(200, res);
        },
        {
          params: t.Object({
            id: t.String({ minLength: 1 }),
          }),
          query: t.Object({
            limit: t.Number({ minimum: 1, default: 15 }),
            last_id: t.Optional(t.Number({ minimum: 1 })),
          }),
        },
      );
  });
}
