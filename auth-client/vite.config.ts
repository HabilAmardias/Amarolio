import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), "");
  const siteUrl = (
    env.VITE_AUTH_CLIENT_URL || "https://auth.amarolio.id"
  ).replace(/\/+$/, "");

  return {
    plugins: [
      react(),
      {
        name: "inject-auth-client-url",
        transformIndexHtml(html) {
          return html.replaceAll("%VITE_AUTH_CLIENT_URL%", siteUrl);
        },
      },
    ],
  };
});
