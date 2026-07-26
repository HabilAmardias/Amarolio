import type { ServerResponse } from "../models/type";
import type { User, LogoutRes, LoginRes } from "../models/user/type";
import { apiFetch } from "./api";

export async function login(): Promise<void> {
  const redirectURI = window.location.origin;
  const reqBody = JSON.stringify({
    redirect_uri: redirectURI
  })
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/login`;
  const res = await apiFetch(url, 200, {
    body: reqBody,
    method: "POST",
  });
  const resBody: ServerResponse<LoginRes> = await res.json();
  window.location.href = resBody.data.redirect_uri;
}

export async function logout(): Promise<void> {
  const redirectURI = window.location.origin;
  const reqBody = JSON.stringify({
    redirect_uri: redirectURI
  })
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/logout`;
  const res = await apiFetch(url, 200, {
    body: reqBody,
    method: "POST",
  });
  const resBody: ServerResponse<LogoutRes> = await res.json();
  window.location.href = resBody.data.redirect_uri;
}

export async function getMe(): Promise<User | null> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/me`;
  const res = await apiFetch(url, 200, {
    method: "GET",
  });
  const resBody: ServerResponse<User> = await res.json();
  return resBody.data;
}
