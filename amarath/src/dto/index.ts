export type ServerResponse<T> = {
  success: boolean;
  data: T;
};

export type TextResponse = {
  message: string;
};
