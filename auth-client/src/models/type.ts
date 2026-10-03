export interface ServerResponse<T> {
  success: boolean;
  data: T;
}

export interface ErrorResponse {
  detail: string;
  error_code: number;
}

export class ApiError extends Error {
  code?: number;
  status?: number;

  constructor(message: string, code?: number, status?: number) {
    super(message);
    this.name = "ApiError";
    this.code = code;
    this.status = status;
  }
}
