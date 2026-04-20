export interface User {
  id: string;
  email: string;
  name: string;
}

// Jotai atoms
import { atom } from "jotai";

export const authAtom = atom<User | null>(null);
export const authLoadingAtom = atom<boolean>(true); // true while restoring session
