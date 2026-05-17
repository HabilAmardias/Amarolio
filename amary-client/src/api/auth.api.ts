import type { ServerResponse } from "../models/model";
import type { User } from "../models/user.model";
import { apiFetch } from "./api";

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

export async function getMe(): Promise<User | null> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/me`;
  const res = await apiFetch(url, {
    method: "GET",
  });
  if (!res.ok) {
    throw new Error("Cannot get user");
  }
  const resBody: ServerResponse<User> = await res.json();
  return resBody.data;
}
