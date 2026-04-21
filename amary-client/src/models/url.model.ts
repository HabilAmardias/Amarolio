import { atom } from "jotai";

export interface ShortenRequest {
  originalUrl: string;
  expiresInDays: number | null; // null = no expiration (auth only)
}

export interface ShortenResponse {
  shortUrl: string;
  originalUrl: string;
  expiresAt: string | null; // ISO date string or null
  createdAt: string;
}

export const urlHistoryAtom = atom<ShortenResponse[]>([]);
