import { GetMessagesParam, SendMessageParam } from "../entity/MessageEntity";
import { Message } from "../entity/MessageEntity";
import { ServerResponse } from "../dto";
import {
  GetMessagesRes,
  SendMessageReq,
  SendMessageRes,
} from "../dto/MessageDto";
import Elysia, { t } from "elysia";
import { CustomError, UnauthorizedError } from "../customerror";

interface MessageServiceItf {
  SendMessage: (param: SendMessageParam) => Promise<string>;
  GetMessages: (param: GetMessagesParam) => Promise<Message[]>;
}

export function NewMessageController(messageService: MessageServiceItf) {
  return new Elysia().group("/api/v1", (app) => {
    return app
      .guard({
        cookie: t.Object({
          user_id: t.Optional(t.String({ minLength: 1 })),
        }),
      })
      .resolve(({ cookie }) => {
        if (!cookie.user_id.value) {
          throw new CustomError(
            "unauthorized access",
            UnauthorizedError,
            "user id not found",
          );
        }
        return { userID: cookie.user_id.value };
      })
      .get(
        "/chatrooms/:id/messages",
        async ({ userID, params, query, status }) => {
          const param = {
            userID,
            chatroomID: params.id,
            limit: query.limit,
            lastID: query.lastID,
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
            lastID: t.Optional(t.Number({ minimum: 1 })),
          }),
        },
      )
      .post(
        "/chatrooms/:id/messages",
        async ({ userID, body, params, status }) => {
          const param = {
            id: params.id,
            userID,
            userMessage: body.user_message,
          };
          const message = await messageService.SendMessage(param);
          const res: ServerResponse<SendMessageRes> = {
            success: true,
            data: {
              message,
            },
          };
          return status(201, res);
        },
        {
          params: t.Object({
            id: t.String({ minLength: 1 }),
          }),
          body: SendMessageReq,
        },
      );
  });
}
