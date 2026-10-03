import { useAtom } from "jotai";
import { userModel } from "../models/user/model";
import { getMe as getMeApi } from "../api/auth.api";
import { useEffect, useCallback, useState } from "react";
import { redirectToLogin, redirectToLogout } from "../lib/authRedirect";

export function useAuth() {
  const [user, setUser] = useAtom(userModel.userAtom);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    getMeApi()
      .then((userData) => setUser(userData))
      .catch(() => setUser(null))
      .finally(() => setIsLoading(false));
  }, [setUser, setIsLoading]);

  const login = useCallback(() => {
    redirectToLogin();
  }, []);

  const logout = useCallback(() => {
    redirectToLogout();
  }, []);

  return { user, isLoading, login, logout };
}
