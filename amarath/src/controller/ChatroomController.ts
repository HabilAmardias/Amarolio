import Elysia, { t } from "elysia";
import { ServerResponse, TextResponse } from "../dto";
import { CreateChatroomRes, GetChatroomRes } from "../dto/ChatroomDto";
import { CustomError, UnauthorizedError } from "../customerror";
import {
  Chatroom,
  CreateChatroomParam,
  DeleteChatroomParam,
  GetChatroomsParam,
} from "../entity/ChatroomEntity";

interface ChatroomServiceItf {
  GetChatrooms: (param: GetChatroomsParam) => Promise<Chatroom[]>;
  DeleteChatroom: (param: DeleteChatroomParam) => Promise<void>;
  CreateChatroom: (param: CreateChatroomParam) => Promise<string>;
}

export function NewChatroomController(chatroomService: ChatroomServiceItf) {
  return new Elysia().group("/api/v1", (app) => {
    return app
      .guard({
        cookie: t.Object({
          user_id: t.Optional(t.String()),
        }),
      })
      .resolve(({ cookie }) => {
        if (!cookie.user_id.value) {
          throw new CustomError(
            "unauthorized access",
            UnauthorizedError,
            "no user id found",
          );
        }
        return { userID: cookie.user_id.value };
      })
      .post(
        "/chatrooms",
        async ({ userID, body, status }) => {
          const param = {
            userID,
            userQuery: body.user_message,
          };
          const id = await chatroomService.CreateChatroom(param);
          const res: ServerResponse<CreateChatroomRes> = {
            success: true,
            data: {
              id,
            },
          };
          return status(201, res);
        },
        {
          body: t.Object({
            user_message: t.String({ minLength: 1 }),
          }),
        },
      )
      .get(
        "/chatrooms",
        async ({ query, status, userID }) => {
          const param = {
            userID,
            limit: query.limit,
            page: query.page,
          };
          const chatrooms = await chatroomService.GetChatrooms(param);
          const res: ServerResponse<GetChatroomRes> = {
            success: true,
            data: {
              chatrooms,
            },
          };
          return status(200, res);
        },
        {
          query: t.Object({
            page: t.Number({ default: 1, minimum: 1 }),
            limit: t.Number({ default: 15, minimum: 1 }),
          }),
        },
      )
      .delete(
        "/chatrooms/:id",
        async ({ userID, params, status }) => {
          const param = {
            userID,
            chatroomID: params.id,
          };
          await chatroomService.DeleteChatroom(param);

          const res: ServerResponse<TextResponse> = {
            success: true,
            data: {
              message: "success delete chatroom",
            },
          };

          return status(200, res);
        },
        {
          params: t.Object({
            id: t.String({ minLength: 1 }),
          }),
        },
      );
  });
}
