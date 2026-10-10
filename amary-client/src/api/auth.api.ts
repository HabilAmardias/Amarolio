import type { ServerResponse } from "../models/type";
import type { User } from "../models/user/type";
import { apiFetch } from "./api";

export async function getMe(): Promise<User | null> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/me`;
  const res = await apiFetch(url, 200, {
    method: "GET",
  });
  const resBody: ServerResponse<User> = await res.json();
  return resBody.data;
}
