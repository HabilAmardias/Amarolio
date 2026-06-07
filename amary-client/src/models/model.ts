export interface ServerResponse<T> {
  success: boolean;
  data: T;
}

export interface PaginateResponse<T> {
  entries: T[];
  page_info: {
    total_row: number;
    last_id?: number;
    page?: number;
    limit: number;
    filter_by?: { name: string; value: unknown }[];
    sort_by?: { name: string; ascend: boolean }[];
  };
}
