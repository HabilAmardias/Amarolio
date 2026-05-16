import type { ServerResponse } from "../models/model";
import type { User } from "../models/user.model";

export async function login(): Promise<void> {
  const redirectURI = window.location.origin;
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/login?redirect_uri=${redirectURI}`;
  window.location.href = url;
}

export async function logout(): Promise<void> {
  const redirectURI = window.location.origin;
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/logout?redirect_uri=${redirectURI}`;
  window.location.href = url;
}

export async function refreshAuth() {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/refresh`;
  const res = await fetch(url, {
    method: "POST",
    credentials: "include",
  });
  if (!res.ok) {
    throw new Error("Auth Expired");
  }
}

export async function getMe(): Promise<User | null> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/me`;
  const res = await fetch(url, {
    method: "GET",
    credentials: "include",
  });
  if (!res.ok) {
    throw new Error("Cannot get user");
  }
  const resBody: ServerResponse<User> = await res.json();
  return resBody.data;
}
