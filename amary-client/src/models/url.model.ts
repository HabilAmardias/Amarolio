import { atom } from "jotai";

export interface FindCustomURLRequest {
  custom_code: string;
}

export interface ShortenRequest {
  url: string;
  duration: number | null; // null = no expiration (auth only)
  custom_code: string | null;
}

export interface ShortenResponse {
  url: string;
  original_url: string;
  expired_at: string | null;
}

export interface UserLink {
  id: number;
  user_id: string | null;
  short_url: string;
  url: string;
  created_at: string;
  expired_at: string | null;
}

const initialToken: string = "";
export const tokenAtom = atom<string>(initialToken);

const initialResult: ShortenResponse | null = null;
export const resultAtom = atom<ShortenResponse | null>(initialResult);
