import { atom } from "jotai";

export interface User {
  username: string;
}

export interface LogoutRes {
  redirect_uri: string;
}

export interface LoginRes {
  redirect_uri: string;
}

const initialUser: User | null = null;
export const authAtom = atom<User | null>(initialUser);
export const authLoadingAtom = atom<boolean>(true); // true while restoring session
