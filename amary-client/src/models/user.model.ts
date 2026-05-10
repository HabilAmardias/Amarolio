import { atom } from "jotai";

export interface User {
  id: string;
  email: string;
  name: string;
}

const initialUser: User | null = null;
export const authAtom = atom<User | null>(initialUser);
export const authLoadingAtom = atom<boolean>(true); // true while restoring session
