// MOCK — replace function bodies with real HTTP calls
import type { ShortenRequest, ShortenResponse } from "../models/url.model";
import type { ServerResponse } from "../models/model";

export async function shortenUrl(
  req: ShortenRequest,
): Promise<ShortenResponse> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url`;
  const res = await fetch(url, {
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
