import { ApiError } from "../models/type";

export function toApiError(err: unknown): ApiError {
  if (err instanceof ApiError) return err;
  return new ApiError(
    err instanceof Error ? err.message : "Something went wrong. Please try again."
  );
}
