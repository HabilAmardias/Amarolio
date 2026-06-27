import type {
  ShortenRequest,
  ShortenResponse,
  UserLink,
  FindCustomURLRequest,
  VisitDashboardRes,
} from "../models/url.model";
import type { ServerResponse, PaginateResponse } from "../models/model";
import { apiFetch } from "./api";

export async function findCustomURLs(customCode: string) {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url/custom-code`;
  const body: FindCustomURLRequest = {
    custom_code: customCode,
  };
  await apiFetch(url, 200, {
    method: "POST",
    body: JSON.stringify(body),
  });
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

  const res = await apiFetch(url, 200, {
    method: "GET",
  });

  const resBody: ServerResponse<PaginateResponse<UserLink>> = await res.json();
  return resBody.data;
}

export async function shortenUrl(
  token: string,
  req: ShortenRequest,
): Promise<ShortenResponse> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url?token=${encodeURIComponent(token)}`;
  const res = await apiFetch(url, 201, {
    method: "POST",
    body: JSON.stringify(req),
  });
  const resBody: ServerResponse<ShortenResponse> = await res.json();
  return resBody.data;
}

export async function getURLMetadata(urlId: string) {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url/${urlId}/metadata`
  const res = await apiFetch(url, 200, {
    method: "GET"
  })
  const resBody: ServerResponse<{ url: UserLink }> = await res.json()
  return resBody.data
}

export async function getURLStats(urlId: string): Promise<VisitDashboardRes> {
  const url = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url/${urlId}/dashboard`
  const res = await apiFetch(url, 200, {
    method: "GET",
  })
  const resBody: ServerResponse<VisitDashboardRes> = await res.json();
  return resBody.data
}

