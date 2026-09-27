import type { ErrorResponse, ServerResponse } from "../models/type";
import { ApiError } from "../models/type";

const BASE_URL = (
  (import.meta.env.VITE_SERVER_HOST as string | undefined) ?? ""
).replace(/\/+$/, "");

export async function apiFetch<T>(
  path: string,
  expectedStatus: number,
  init?: RequestInit,
): Promise<ServerResponse<T>> {
  const headers: Record<string, string> = {};
  if (init?.body) {
    headers["Content-Type"] = "application/json";
  }

  const res = await fetch(`${BASE_URL}${path}`, {
    ...init,
    credentials: "include",
    headers: { ...headers, ...(init?.headers as Record<string, string>) },
  });

  if (res.status === expectedStatus) {
    return (await res.json()) as ServerResponse<T>;
  }

  let detail = `HTTP Error ${res.status}`;
  let code: number | undefined;

  const contentType = res.headers.get("content-type");
  if (contentType?.includes("application/json")) {
    try {
      const body = (await res.json()) as ServerResponse<ErrorResponse>;
      detail = body?.data?.detail || detail;
      code = body?.data?.error_code;
    } catch {
      // keep the fallback detail
    }
  }

  throw new ApiError(detail, code, res.status);
}
