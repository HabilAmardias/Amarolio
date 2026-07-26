import { useAtom } from "jotai";
import { userModel } from "../models/user/model";
import {
  login as loginApi,
  logout as logoutApi,
  getMe as getMeApi,
} from "../api/auth.api";
import { useEffect, useCallback, useState } from "react";

export function useAuth() {
  const [user, setUser] = useAtom(userModel.userAtom);
  const [isLoading, setIsLoading] = useState(true);

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
