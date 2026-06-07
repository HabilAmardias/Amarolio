import { useState } from "react";
import { useAtom } from "jotai";
import { authAtom } from "../models/user.model";
import { tokenAtom, resultAtom } from "../models/url.model";
import { shortenUrl } from "../api/url.api";

export function useShorten() {
  const [user] = useAtom(authAtom);
  const [result, setResult] = useAtom(resultAtom);
  const [token, setToken] = useAtom(tokenAtom);

  const [url, setUrl] = useState("");
  const [expiresInDays, setExpiresInDays] = useState<number | null>(1);
  const [noExpiry, setNoExpiry] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleShorten = async () => {
    setError(null);
    setResult(null);

    try {
      new URL(url);
    } catch {
      setError("Please enter a valid URL");
      return;
    }

    const finalExpiresInDays = user ? (noExpiry ? null : expiresInDays) : 1;

    setIsLoading(true);
    if (!token) {
      setError("invalid request");
      return;
    }
    try {
      const response = await shortenUrl(token, {
        url: url,
        duration: finalExpiresInDays,
      });
      setResult(response);
      setUrl("");
    } catch (err) {
      setError((err as Error).message);
    } finally {
      setIsLoading(false);
    }
  };

  return {
    url,
    setUrl,
    expiresInDays,
    setExpiresInDays,
    noExpiry,
    setNoExpiry,
    token,
    setToken,
    result,
    error,
    isLoading,
    handleShorten,
  };
}
