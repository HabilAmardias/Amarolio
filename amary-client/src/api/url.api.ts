import type { ShortenRequest, ShortenResponse } from "../models/url.model";
import type { ServerResponse } from "../models/model";
import { apiFetch } from "./api";

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

export async function getUserUrls(): Promise<ShortenResponse[]> {
  // MOCK: return empty history
  return [];
}
