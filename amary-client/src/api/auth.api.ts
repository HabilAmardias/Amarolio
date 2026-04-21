// MOCK — replace function bodies with real HTTP calls
import type { User } from "../models/user.model";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export async function login(provider: string, _token: string): Promise<User> {
  // MOCK: simulate successful Google login
  const user: User = {
    id: "google-1",
    email: "user@gmail.com",
    name: "Google User",
  };
  // Store in localStorage for persistence
  localStorage.setItem("auth_user", JSON.stringify(user));
  return user;
}

export async function logout(): Promise<void> {
  // MOCK: clear session cookie/token
  localStorage.removeItem("auth_user");
}

export async function getMe(): Promise<User | null> {
  // MOCK: restore session from localStorage
  const stored = localStorage.getItem("auth_user");
  if (stored) {
    try {
      return JSON.parse(stored) as User;
    } catch {
      return null;
    }
  }
  return null;
}
