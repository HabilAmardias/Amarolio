export interface AppLink {
  id: string;
  name: string;
  description: string;
  url: string;
  accent: string;
}

export const apps: AppLink[] = [
  {
    id: "amary",
    name: "Amary",
    description: "Shorten long links into clean, trackable URLs.",
    url: import.meta.env.VITE_AMARY_CLIENT_URL ?? "https://amary.amarolio.id",
    accent: "#F0A63B",
  },
  {
    id: "amarolio",
    name: "Amarolio",
    description: "Projects, experience, and skills of Muhammad Habil Amardias.",
    url: import.meta.env.VITE_AMAROLIO_CLIENT_URL ?? "https://amarolio.id",
    accent: "#3D6BD4",
  },
];
