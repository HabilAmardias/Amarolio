const STORAGE_KEY = "amarolio.auth.redirect_uri";

export function getAllowedOrigins(): string[] {
  return (
    (import.meta.env.VITE_ALLOWED_REDIRECT_ORIGINS as string | undefined) ?? ""
  )
    .split(";")
    .map((origin) => origin.trim())
    .filter(Boolean);
}

/**
 * Only redirect back to a whitelisted origin. Anything else falls back to the
 * auth-client account page ("/") to avoid open-redirect abuse.
 */
export function resolveRedirectUri(raw: string | null | undefined): string {
  if (raw) {
    try {
      const url = new URL(raw);
      if (getAllowedOrigins().includes(url.origin)) {
        return url.href;
      }
    } catch {
      // not an absolute URL — ignore
    }
  }
  return "/";
}

export function persistRedirectUri(uri: string): void {
  if (uri && uri !== "/") {
    sessionStorage.setItem(STORAGE_KEY, uri);
  } else {
    sessionStorage.removeItem(STORAGE_KEY);
  }
}

export function getPersistedRedirectUri(): string {
  return sessionStorage.getItem(STORAGE_KEY) ?? "/";
}

export function clearPersistedRedirectUri(): void {
  sessionStorage.removeItem(STORAGE_KEY);
}
