import { useAtom } from "jotai";
import { authAtom, authLoadingAtom } from "../models/user.model";
import { login as loginApi, logout as logoutApi, getMe } from "../api/auth.api";
import { useEffect, useCallback } from "react";

export function useAuth() {
  const [user, setUser] = useAtom(authAtom);
  const [isLoading, setIsLoading] = useAtom(authLoadingAtom);

  useEffect(() => {
    // Restore session on mount
    getMe().then((userData) => {
      setUser(userData);
      setIsLoading(false);
    });
  }, [setUser, setIsLoading]);

  const login = useCallback(
    async (provider: string = "google") => {
      if (provider === "google") {
        const userData = await loginApi("google", "");
        setUser(userData);
        // Store in localStorage for persistence
        localStorage.setItem("auth_user", JSON.stringify(userData));
      }
    },
    [setUser],
  );

  const logout = useCallback(async () => {
    await logoutApi();
    setUser(null);
    localStorage.removeItem("auth_user");
  }, [setUser]);

  return { user, isLoading, login, logout };
}
