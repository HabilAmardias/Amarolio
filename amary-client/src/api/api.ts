async function refreshAuth() {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/refresh`;
  const res = await fetch(url, {
    method: "POST",
    credentials: "include",
  });
  if (!res.ok) {
    throw new Error("Auth Expired");
  }
}

export const apiFetch = async (info: RequestInfo, init?: RequestInit) => {
  const res = await fetch(info, {
    ...init,
    credentials: "include",
  });
  if (res.status !== 401) return res;
  await refreshAuth();
  return fetch(info, {
    ...init,
    credentials: "include",
  });
};
