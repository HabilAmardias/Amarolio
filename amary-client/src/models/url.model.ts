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

// Jotai atom
import { atom } from "jotai";

export const urlHistoryAtom = atom<ShortenResponse[]>([]);
