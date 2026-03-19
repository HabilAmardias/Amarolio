import { t } from "elysia";
import { Message } from "../entity/MessageEntity";

export const SendMessageReq = t.Object({
  user_message: t.String({ minLength: 1 }),
});

export type SendMessageRes = {
  message: string;
};

export type GetMessagesRes = {
  messages: Message[];
};
