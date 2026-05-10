import { atom } from "jotai";

export interface ShortenRequest {
  url: string;
  duration: number | null; // null = no expiration (auth only)
}

export interface ShortenResponse {
  url: string;
  original_url: string;
  expired_at: string | null;
}

export const urlHistoryAtom = atom<ShortenResponse[]>([]);

const initialResult: ShortenResponse | null = null;
export const shortenResultAtom = atom(initialResult);
