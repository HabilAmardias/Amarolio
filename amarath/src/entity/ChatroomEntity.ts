export type Chatroom = {
  id: string;
  user_id: string;
  title: string;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};

export type GetChatroomsParam = {
  userID: string;
  limit: number;
  page: number;
};

export type DeleteChatroomParam = {
  userID: string;
  chatroomID: string;
};

export type CreateChatroomParam = {
  userID: string;
  userQuery: string;
};
