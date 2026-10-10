const authClientUrl = (
  (import.meta.env.VITE_AUTH_CLIENT_URL as string | undefined) ?? ""
).replace(/\/+$/, "");

export function redirectToLogin(returnTo?: string): void {
  const target = returnTo ?? window.location.origin;
  window.location.href = `${authClientUrl}/login?redirect_uri=${encodeURIComponent(
    target
  )}`;
}

export function redirectToLogout(): void {
  window.location.href = `${authClientUrl}/logout?redirect_uri=${encodeURIComponent(
    window.location.origin
  )}`;
}
