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

const initialHistory: ShortenResponse[] = [];
export const urlHistoryAtom = atom<ShortenResponse[]>(initialHistory);

const initialToken: string = "";
export const tokenAtom = atom<string>(initialToken);

const initialResult = null;
export const resultAtom = atom<ShortenResponse | null>(initialResult);
