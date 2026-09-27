export const siteConfig = {
  name: "Amarolio Account",
  author: "Muhammad Habil Amardias",
  url: (
    (import.meta.env.VITE_AUTH_CLIENT_URL as string | undefined) ?? ""
  )
    .trim()
    .replace(/\/+$/, "") || "https://auth.amarolio.id",
  description:
    "One secure sign-in for every Amarolio service. Manage your account and access Amary and Amarolio.",
};
