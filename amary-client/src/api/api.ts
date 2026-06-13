import type { ErrorResponse, ServerResponse } from "../models/model";

async function refreshAuth() {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/refresh`;
  const res = await fetch(url, {
    method: "POST",
    credentials: "include",
  });
  if (!res.ok) {
    throw new Error("Session Expired");
  }
}

export const apiFetch = async (
  info: RequestInfo,
  expectedStatus: number,
  init?: RequestInit,
) => {
  let res = await fetch(info, {
    ...init,
    credentials: "include",
  });
  if (res.status === 401) {
    await refreshAuth();

    res = await fetch(info, {
      ...init,
      credentials: "include",
    });

    if (res.status === 401) {
      throw new Error("SESSION_EXPIRED");
    }
  }

  if (res.status !== expectedStatus) {
    let errorMessage = `Request failed with status ${res.status}`;
    const contentType = res.headers.get("content-type");
    if (contentType && contentType.includes("application/json")) {
      try {
        const resBody: ServerResponse<ErrorResponse> = await res.json();
        errorMessage = resBody.data?.detail || errorMessage;
      } catch {
        errorMessage = "Internal Server Error";
      }
    }
    throw new Error(errorMessage);
  }

  return res;
};
