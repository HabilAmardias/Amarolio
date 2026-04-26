import type { Profile, Experience, Project } from "./types";
import AmaryHomepage from "../assets/Amary-Homepage.png";

export const PROFILE: Profile = {
  name: "Amarolio", // {/* TODO: Replace with real content */}
  title: "Portofolio of someone who aspired to be Software Engineer", // {/* TODO: Replace with real content */}
  bio: "Someone who enjoys keeping up with technological advancements", // {/* TODO: Replace with real content */}
  avatarUrl: "/placeholder-avatar.jpg", // {/* TODO: Replace with real content */}
  socials: [
    {
      label: "GitHub",
      url: "https://github.com/HabilAmardias",
      icon: "github",
    },
    {
      label: "LinkedIn",
      url: "https://www.linkedin.com/in/muhammad-habil-amardias/",
      icon: "linkedin",
    },
  ],
};

export const EXPERIENCES: Experience[] = [
  {
    id: "exp-1",
    company: "PT Indomobil Finance", // {/* TODO: Replace with real content */}
    role: "Development System Staff", // {/* TODO: Replace with real content */}
    period: "July 2025 – May 2026", // {/* TODO: Replace with real content */}
    location: "East Jakarta, Indonesia", // {/* TODO: Replace with real content */}
    description:
      "Mediate communication between users and IT team regarding feature development and conduct UAT to test the newly developed features", // {/* TODO: Replace with real content */}
    tags: [],
  },
  {
    id: "exp-2",
    company: "Sealabs", // {/* TODO: Replace with real content */}
    role: "Software Engineer Trainee", // {/* TODO: Replace with real content */}
    period: "Feb 2025 – July 2025", // {/* TODO: Replace with real content */}
    location: "South Jakarta, Indonesia", // {/* TODO: Replace with real content */}
    description: "Full-stack web applications and Web applications deployment", // {/* TODO: Replace with real content */}
    tags: ["React", "Go", "Javascript", "Full-Stack"],
  },
];

export const PROJECTS: Project[] = [
  {
    id: "proj-1",
    title: "Amary", // {/* TODO: Replace with real content */}
    description: "URL Shortener", // {/* TODO: Replace with real content */}
    tags: ["React", "Go", "PostgreSQL", "Redis"],
    imageUrl: AmaryHomepage, // {/* TODO: Replace with real content */}
    liveUrl: import.meta.env.AMARY_CLIENT_DOMAIN || "#",
    repoUrl: "https://github.com/HabilAmardias/Amarolio",
    featured: true,
  },
];
