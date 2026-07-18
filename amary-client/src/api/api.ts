import type { ErrorResponse, ServerResponse } from "../models/model";
import type { LogoutRes } from "../models/user.model";

// Store the active refresh promise to share across concurrent API requests
let activeRefreshPromise: Promise<string> | null = null;

async function refreshAuth(): Promise<string> {
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

      if (!res.ok) {
        let errorDetail = "Internal Server Error";
        let errorCode: number | undefined;

        const contentType = res.headers.get("content-type");
        if (contentType && contentType.includes("application/json")) {
          try {
            const resBody: ServerResponse<ErrorResponse> = await res.json();
            errorCode = resBody?.data?.error_code;
            errorDetail = resBody?.data?.detail || errorDetail;
          } catch (e) {
            // Fallback if JSON parsing fails
          }
        }

        // if other than refresh token expired (or error parsing code not matching), throw an error
        if (errorCode !== 40102) {
          throw new Error(errorDetail);
        }

        // if refresh token expired, remove all cookies and redirect to homepage
        const origin = window.location.origin;
        const reqBody = JSON.stringify({
          redirect_uri: origin
        });
        const logoutURL = `${import.meta.env.VITE_SERVER_HOST}/api/v1/logout`;
        const logoutRes = await fetch(logoutURL, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
          body: reqBody
        });

        if (logoutRes.ok) {
          const contentTypeLogout = logoutRes.headers.get("content-type");
          if (contentTypeLogout && contentTypeLogout.includes("application/json")) {
            try {
              const logoutResBody: ServerResponse<LogoutRes> = await logoutRes.json();
              return logoutResBody.data.redirect_uri;
            } catch (e) {
              // ignore and fallback
            }
          }
        }
        return origin;
      }
      return "";
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
    } catch (e) {
      // Fallback
    }
  }

  const errorCode = resBody?.data?.error_code;
  const errorDetail = resBody?.data?.detail || `HTTP Error ${res.status}`;

  // if the error is about access token expired, then try to refresh the token and retry the request
  if (errorCode === 40102) {
    const redirectURI = await refreshAuth();
    if (redirectURI) {
      window.location.href = redirectURI;
      return new Promise<never>(() => { });
    }

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
      } catch (e) {
        // Fallback
      }
    }
    const retryErrorDetail = retryResBody?.data?.detail || `HTTP Error ${res.status}`;
    throw new Error(retryErrorDetail);
  }

  throw new Error(errorDetail);
};
