import { useAtom } from "jotai";
import { authAtom, authLoadingAtom } from "../models/user.model";
import {
  login as loginApi,
  logout as logoutApi,
  getMe as getMeApi,
} from "../api/auth.api";
import { useEffect, useCallback } from "react";

export function useAuth() {
  const [user, setUser] = useAtom(authAtom);
  const [isLoading, setIsLoading] = useAtom(authLoadingAtom);

  useEffect(() => {
    getMeApi()
      .then((userData) => setUser(userData))
      .catch(() => setUser(null))
      .finally(() => setIsLoading(false));
  }, [setUser, setIsLoading]);

  const login = useCallback(async () => {
    await loginApi();
  }, []);

  const logout = useCallback(async () => {
    await logoutApi();
  }, []);

  return { user, isLoading, login, logout };
}
