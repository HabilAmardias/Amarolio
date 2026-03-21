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

export type SaveMessageQuery = {
  content: string;
  role: string;
};

export interface MessageResponse {
  message: {
    thinking?: string;
    content: string;
  };
  done: boolean;
  done_reason: string;
}

export type GetMessagesParam = {
  userID: string;
  chatroomID: string;
  limit: number;
  lastID?: number;
};
