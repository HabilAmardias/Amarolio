export const siteConfig = {
  name: "Amarolio",
  author: "Muhammad Habil Amardias",
  url: (
    (import.meta.env.VITE_SITE_URL as string | undefined) ?? ""
  )
    .trim()
    .replace(/\/+$/, "") || "https://amarolio.id",
  description:
    "Muhammad Habil Amardias — software engineer. Explore projects, professional experience, and skills.",
  locale: "en_US",
  image: "/og-image.jpg",
};
