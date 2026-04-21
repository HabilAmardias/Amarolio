// MOCK — replace function bodies with real HTTP calls
import type { ShortenRequest, ShortenResponse } from "../models/url.model";

export async function shortenUrl(
  req: ShortenRequest,
): Promise<ShortenResponse> {
  // MOCK: simulate shortening
  return {
    shortUrl: `https://short.ly/${Math.random().toString(36).slice(2, 8)}`,
    originalUrl: req.originalUrl,
    expiresAt: req.expiresInDays
      ? new Date(Date.now() + req.expiresInDays * 86400000).toISOString()
      : null,
    createdAt: new Date().toISOString(),
  };
}

export async function getUserUrls(): Promise<ShortenResponse[]> {
  // MOCK: return empty history
  return [];
}
