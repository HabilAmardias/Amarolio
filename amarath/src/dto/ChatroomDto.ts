export type ChatroomDto = {
  id: string;
  user_id: string;
  title: string;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
};

export type CreateChatroomReq = {
  user_query: string;
};

export type GetChatroomRes = {
  chatrooms: ChatroomDto[];
};

export type CreateChatroomRes = {
  id: string;
};
