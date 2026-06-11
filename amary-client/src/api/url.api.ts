import type {
  ShortenRequest,
  ShortenResponse,
  UserLink,
} from "../models/url.model";
import type { ServerResponse, PaginateResponse } from "../models/model";
import { apiFetch } from "./api";

export async function findCustomURLs(customCode: string) {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url/find/${customCode}`;
  const res = await apiFetch(url, {
    method: "GET",
  });
  if (!res.ok) {
    throw new Error("url already exist");
  }
  return true;
}

export async function getUserURLs(limit?: number, lastID?: number) {
  let url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/me/url?`;
  const args: string[] = [];

  if (limit) {
    args.push(`limit=${limit}`);
  }
  if (lastID) {
    args.push(`last_id=${lastID}`);
  }
  url += args.join("&");
  const res = await apiFetch(url, {
    method: "GET",
  });
  if (!res.ok) {
    throw new Error("something went wrong");
  }

  const resBody: ServerResponse<PaginateResponse<UserLink>> = await res.json();
  return resBody.data;
}

export async function shortenUrl(
  token: string,
  req: ShortenRequest,
): Promise<ShortenResponse> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url?token=${encodeURIComponent(token)}`;
  const res = await apiFetch(url, {
    method: "POST",
    body: JSON.stringify(req),
  });
  const resBody: ServerResponse<ShortenResponse> = await res.json();
  return resBody.data;
}
