export type Message = {
  id: number;
  chatroom_id: string;
  role: string;
  content: string;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};

export type SendMessageParam = {
  id: string;
  userID: string;
  userMessage: string;
};

export type GetMessagesParam = {
  userID: string;
  chatroomID: string;
  limit: number;
  lastID?: number;
};
