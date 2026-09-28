import { apiFetch } from "./apiFetch";
import { ApiError } from "../models/type";
import type { User } from "../models/user/type";

export async function prelogin(
  email: string,
  password: string,
): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/prelogin", 200, {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}

export async function login(otp: string): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/login", 200, {
    method: "POST",
    body: JSON.stringify({ otp }),
  });
}

export async function resendOtp(): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/otp/send", 200, {
    method: "GET",
  });
}

export async function register(
  email: string,
  password: string,
): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/register", 200, {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}

export async function resendVerification(email: string): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/verify/send", 200, {
    method: "POST",
    body: JSON.stringify({ email }),
  });
}

export async function verify(userId: string, token: string): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/verify", 200, {
    method: "POST",
    body: JSON.stringify({ user_id: userId, token }),
  });
}

export async function sendResetPasswordEmail(email: string): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/reset-password/send", 200, {
    method: "POST",
    body: JSON.stringify({ email }),
  });
}

export async function resetPassword(
  userId: string,
  token: string,
  newPassword: string,
): Promise<void> {
  await apiFetch<{ message: string }>("/api/v1/reset-password", 200, {
    method: "POST",
    body: JSON.stringify({
      user_id: userId,
      token,
      new_password: newPassword,
    }),
  });
}

export async function getMe(): Promise<User | null> {
  try {
    const body = await apiFetch<User>("/api/v1/me", 200, { method: "GET" });
    return body.data;
  } catch (err) {
    if (err instanceof ApiError && (err.status === 401 || err.code === 40101)) {
      return null;
    }
    throw err;
  }
}

export async function logout(): Promise<void> {
  try {
    await apiFetch<{ message: string }>("/api/v1/logout", 200, {
      method: "POST",
    });
  } catch {
    // Logging out is best effort — the cookies are cleared server side.
  }
}
