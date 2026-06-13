export interface ServerResponse<T> {
  success: boolean;
  data: T;
}

export interface PaginateResponse<T> {
  entries: T[];
  page_info: {
    last_id?: number;
    page?: number;
    limit: number;
    filter_by?: { name: string; value: unknown }[];
    sort_by?: { name: string; ascend: boolean }[];
  };
}

export interface ErrorResponse {
  detail: string;
}
