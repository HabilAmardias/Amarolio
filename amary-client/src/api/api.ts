import type { ErrorResponse, ServerResponse } from "../models/type";
import { redirectToLogin } from "../lib/authRedirect";

// Store the active refresh promise to share across concurrent API requests
let activeRefreshPromise: Promise<void> | null = null;

async function refreshAuth(): Promise<void> {
  if (activeRefreshPromise) {
    return activeRefreshPromise;
  }

  activeRefreshPromise = (async () => {
    try {
      const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/refresh`;
      const res = await fetch(url, {
        method: "POST",
        credentials: "include",
      });

      if (res.ok) {
        return;
      }

      let errorDetail = "Internal Server Error";
      let errorCode: number | undefined;

      const contentType = res.headers.get("content-type");
      if (contentType && contentType.includes("application/json")) {
        try {
          const resBody: ServerResponse<ErrorResponse> = await res.json();
          errorCode = resBody?.data?.error_code;
          errorDetail = resBody?.data?.detail || errorDetail;
        } catch {
          // Fallback if JSON parsing fails
        }
      }

      // if the refresh token is still valid, surface the original error
      if (errorCode !== 40102) {
        throw new Error(errorDetail);
      }

      // refresh token expired — hand off to the centralized auth client
      redirectToLogin(window.location.origin);
      await new Promise<never>(() => {});
    } finally {
      activeRefreshPromise = null;
    }
  })();

  return activeRefreshPromise;
}

export const apiFetch = async (
  info: RequestInfo,
  expectedStatus: number,
  init?: RequestInit,
): Promise<Response> => {
  let res = await fetch(info, {
    ...init,
    credentials: "include",
  });

  // if http status match with expected status, then return response
  if (res.status === expectedStatus) {
    return res;
  }

  // Parse error details safely
  let resBody: ServerResponse<ErrorResponse> | null = null;
  const contentType = res.headers.get("content-type");
  if (contentType && contentType.includes("application/json")) {
    try {
      resBody = await res.json();
    } catch {
      // Fallback
    }
  }

  const errorCode = resBody?.data?.error_code;
  const errorDetail = resBody?.data?.detail || `HTTP Error ${res.status}`;

  // if the error is about access token expired, refresh the token and retry
  if (errorCode === 40102) {
    await refreshAuth();

    res = await fetch(info, {
      ...init,
      credentials: "include",
    });

    if (res.status === expectedStatus) {
      return res;
    }

    let retryResBody: ServerResponse<ErrorResponse> | null = null;
    const retryContentType = res.headers.get("content-type");
    if (retryContentType && retryContentType.includes("application/json")) {
      try {
        retryResBody = await res.json();
      } catch {
        // Fallback
      }
    }
    const retryErrorDetail =
      retryResBody?.data?.detail || `HTTP Error ${res.status}`;
    throw new Error(retryErrorDetail);
  }

  throw new Error(errorDetail);
};
