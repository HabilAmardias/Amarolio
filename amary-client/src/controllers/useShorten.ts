import { useState } from "react";
import { useAtom } from "jotai";
import { authAtom } from "../models/user.model";
import { urlHistoryAtom, shortenResultAtom } from "../models/url.model";
import { shortenUrl } from "../api/url.api";

export function useShorten() {
  const [user] = useAtom(authAtom);
  const [, setUrlHistory] = useAtom(urlHistoryAtom);
  const [result, setResult] = useAtom(shortenResultAtom);

  const [url, setUrl] = useState("");
  const [expiresInDays, setExpiresInDays] = useState<number | null>(1);
  const [noExpiry, setNoExpiry] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleShorten = async () => {
    setError(null);
    setResult(null);

    // Validate URL
    try {
      new URL(url);
    } catch {
      setError("Please enter a valid URL");
      return;
    }

    // If unauthenticated, force 1 day expiration
    const finalExpiresInDays = user ? (noExpiry ? null : expiresInDays) : 1;

    setIsLoading(true);
    try {
      const response = await shortenUrl({
        url: url,
        duration: finalExpiresInDays,
      });
      setResult(response);
      setUrlHistory((prev) => [response, ...prev]);
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
    result,
    error,
    isLoading,
    handleShorten,
  };
}
