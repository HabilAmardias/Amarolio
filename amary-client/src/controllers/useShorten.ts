import { useState, useRef, useEffect } from "react";
import { useAtom } from "jotai";
import { userModel } from "../models/user/model";
import { shortenUrl, findCustomURLs } from "../api/url.api";
import type { ShortenResponse } from "../models/url/type";

export function useShorten() {
  const [user] = useAtom(userModel.userAtom);
  const [result, setResult] = useState<ShortenResponse | null>(null);
  const [token, setToken] = useState<string>("");
  const [rawCustom, setRawCustom] = useState<string>("");
  const [customSlug, setCustomSlug] = useState<string>("");
  const timerRef = useRef<number | null>(null);

  const [url, setUrl] = useState("");
  const [expiresInDays, setExpiresInDays] = useState<number | null>(1);
  const [noExpiry, setNoExpiry] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [isCheckingSlug, setIsCheckingSlug] = useState(false);

  useEffect(() => {
    return () => {
      if (timerRef.current) {
        clearTimeout(timerRef.current);
      }
    };
  }, []);

  const onCustomChange = (value: string) => {
    setRawCustom(value);
    if (timerRef.current) {
      clearTimeout(timerRef.current);
    }
    // start debounce
    timerRef.current = window.setTimeout(async () => {
      const trimmed = value.trim();
      setCustomSlug(trimmed);
      timerRef.current = null;
      setError(null);

      if (!trimmed) {
        setIsCheckingSlug(false);
        return;
      }

      setIsCheckingSlug(true);
      try {
        await findCustomURLs(trimmed);
      } catch (e) {
        setError((e as Error).message);
      } finally {
        setIsCheckingSlug(false);
      }
    }, 500);
  };

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

    if (!token) {
      setError("invalid request");
      return;
    }

    setIsLoading(true);
    try {
      const customCode = customSlug || (rawCustom ? rawCustom.trim() : null);
      const response = await shortenUrl(token, {
        url: url,
        duration: finalExpiresInDays,
        custom_code: customCode,
      });

      setResult(response);
      setUrl("");
      setRawCustom("");
      setCustomSlug("");
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
    rawCustom,
    onCustomChange,
    error,
    isLoading,
    isCheckingSlug,
    handleShorten,
  };
}
