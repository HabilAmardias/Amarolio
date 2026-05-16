import { atom } from "jotai";

export interface User {
  email: string;
}

const initialUser: User | null = null;
export const authAtom = atom<User | null>(initialUser);
export const authLoadingAtom = atom<boolean>(true); // true while restoring session
