Project Overview
Build a personal portfolio web app for a developer/designer. All personal content (name, bio, experience, projects) should use clearly labeled placeholder text. The app must score ≥90 on Lighthouse (Performance, Accessibility, Best Practices, SEO).

Tech Stack (no additional dependencies)

Vite — build tool
React + TypeScript — UI framework
React Router v6 — client-side routing
Jotai — global state management
Material UI (MUI) v5 — component library & theming


Architecture: MVC + Folder-Based Pattern
src/
├── models/               # Data shapes & mock data
│   ├── types.ts          # All TypeScript interfaces (Experience, Project, Profile, etc.)
│   └── data.ts           # Placeholder data arrays exported as constants
│
├── controllers/          # Business logic, state orchestration
│   ├── useProfileController.ts
│   ├── useExperienceController.ts
│   └── useProjectController.ts
│
├── views/                # Pages and page-level components
│   ├── Home/
│   │   ├── index.tsx
│   │   └── HomeView.tsx
│   ├── Experience/
│   │   ├── index.tsx
│   │   └── ExperienceView.tsx
│   ├── Projects/
│   │   ├── index.tsx
│   │   └── ProjectsView.tsx
│   └── Contact/
│       ├── index.tsx
│       └── ContactView.tsx
│
├── components/           # Reusable UI components
│   ├── layout/
│   │   ├── Navbar.tsx
│   │   └── Footer.tsx
│   ├── cards/
│   │   ├── ProjectCard.tsx
│   │   └── ExperienceCard.tsx
│   └── common/
│       ├── SectionHeader.tsx
│       └── GlassBox.tsx    # Core glass UI wrapper
│
├── store/                # Jotai atoms
│   └── atoms.ts
│
├── theme/                # MUI theme config
│   └── theme.ts
│
├── router/               # Route definitions
│   └── index.tsx
│
└── main.tsx
Convention rules:

models/ = pure data — no React, no side effects
controllers/ = custom hooks that read from atoms/models and expose data + handlers to views
views/ = pages composed of components, wired to controllers via hooks
components/ = stateless or lightly stateful UI primitives, no direct data fetching


Models (src/models/types.ts)
tsexport interface Profile {
  name: string;
  title: string;
  bio: string;
  avatarUrl: string;
  socials: { label: string; url: string; icon: string }[];
}

export interface Experience {
  id: string;
  company: string;
  role: string;
  period: string;        // e.g. "Jan 2022 – Present"
  location: string;
  description: string;
  tags: string[];        // tech stack used
}

export interface Project {
  id: string;
  title: string;
  description: string;
  tags: string[];
  imageUrl: string;
  liveUrl?: string;
  repoUrl?: string;
  featured: boolean;
}

Placeholder Data (src/models/data.ts)
Populate with clearly labeled placeholders, e.g.:
tsexport const PROFILE: Profile = {
  name: "Your Name",
  title: "Full-Stack Developer & Designer",
  bio: "A short bio about yourself goes here. Highlight what makes you unique...",
  avatarUrl: "/placeholder-avatar.jpg",
  socials: [
    { label: "GitHub", url: "#", icon: "github" },
    { label: "LinkedIn", url: "#", icon: "linkedin" },
  ],
};

export const EXPERIENCES: Experience[] = [
  {
    id: "exp-1",
    company: "Company Name",
    role: "Job Title",
    period: "Month Year – Present",
    location: "City, Country",
    description: "Describe your responsibilities and achievements here...",
    tags: ["React", "TypeScript", "Node.js"],
  },
  // Add 2–3 more placeholders
];

export const PROJECTS: Project[] = [
  {
    id: "proj-1",
    title: "Project Title",
    description: "What does this project do? What problem does it solve?",
    tags: ["React", "MUI", "Vite"],
    imageUrl: "/placeholder-project.jpg",
    liveUrl: "#",
    repoUrl: "#",
    featured: true,
  },
  // Add 3–5 more placeholders
];

Jotai Store (src/store/atoms.ts)
tsimport { atom } from "jotai";
import { Profile, Experience, Project } from "../models/types";
import { PROFILE, EXPERIENCES, PROJECTS } from "../models/data";

export const profileAtom = atom<Profile>(PROFILE);
export const experiencesAtom = atom<Experience[]>(EXPERIENCES);
export const projectsAtom = atom<Project[]>(PROJECTS);
export const activeFilterAtom = atom<string>("all"); // for project tag filtering

Controllers (example)
ts// src/controllers/useProjectController.ts
import { useAtom } from "jotai";
import { projectsAtom, activeFilterAtom } from "../store/atoms";

export function useProjectController() {
  const [projects] = useAtom(projectsAtom);
  const [activeFilter, setActiveFilter] = useAtom(activeFilterAtom);

  const allTags = ["all", ...Array.from(new Set(projects.flatMap(p => p.tags)))];

  const filtered = activeFilter === "all"
    ? projects
    : projects.filter(p => p.tags.includes(activeFilter));

  return { projects: filtered, allTags, activeFilter, setActiveFilter };
}

Routing (src/router/index.tsx)
tsximport { createBrowserRouter, RouterProvider } from "react-router-dom";
import HomeView from "../views/Home";
import ExperienceView from "../views/Experience";
import ProjectsView from "../views/Projects";
import ContactView from "../views/Contact";
import RootLayout from "../components/layout/RootLayout";

const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      { index: true, element: <HomeView /> },
      { path: "experience", element: <ExperienceView /> },
      { path: "projects", element: <ProjectsView /> },
      { path: "contact", element: <ContactView /> },
    ],
  },
]);

export default function AppRouter() {
  return <RouterProvider router={router} />;
}

Visual Direction: Glass Theme
Concept: Deep dark background (near-black navy or charcoal) with frosted-glass panels floating above it. Subtle chromatic light leaks and soft glows for depth. MUI components are overridden to match.
MUI Theme (src/theme/theme.ts) key config:
tsimport { createTheme, alpha } from "@mui/material/styles";

export const theme = createTheme({
  palette: {
    mode: "dark",
    background: { default: "#090d1a", paper: "rgba(255,255,255,0.04)" },
    primary: { main: "#7eb8f7" },
    secondary: { main: "#a78bfa" },
    text: { primary: "#e8eaf6", secondary: "#8892b0" },
  },
  typography: {
    fontFamily: "'Sora', 'DM Sans', sans-serif",
    h1: { fontWeight: 700, letterSpacing: "-0.03em" },
  },
  components: {
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundImage: "none",
          backdropFilter: "blur(16px) saturate(180%)",
          background: "rgba(255, 255, 255, 0.04)",
          border: "1px solid rgba(255, 255, 255, 0.08)",
          borderRadius: 16,
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          backdropFilter: "blur(20px)",
          background: "rgba(255,255,255,0.05)",
          border: "1px solid rgba(255,255,255,0.1)",
          transition: "transform 0.2s ease, box-shadow 0.2s ease",
          "&:hover": {
            transform: "translateY(-4px)",
            boxShadow: "0 20px 60px rgba(0,0,0,0.4)",
          },
        },
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          background: "rgba(9, 13, 26, 0.7)",
          backdropFilter: "blur(20px)",
          borderBottom: "1px solid rgba(255,255,255,0.06)",
          boxShadow: "none",
        },
      },
    },
  },
});
Background effect: Add a global CSS mesh gradient behind everything:
cssbody::before {
  content: "";
  position: fixed;
  inset: 0;
  z-index: -1;
  background:
    radial-gradient(ellipse at 20% 50%, rgba(126,184,247,0.08) 0%, transparent 60%),
    radial-gradient(ellipse at 80% 20%, rgba(167,139,250,0.08) 0%, transparent 50%),
    #090d1a;
}
GlassBox component — wrap any section in this:
tsximport { Box, BoxProps, alpha } from "@mui/material";

export default function GlassBox({ children, sx, ...props }: BoxProps) {
  return (
    <Box
      sx={{
        backdropFilter: "blur(20px) saturate(160%)",
        background: "rgba(255,255,255,0.04)",
        border: "1px solid rgba(255,255,255,0.09)",
        borderRadius: 4,
        p: 4,
        ...sx,
      }}
      {...props}
    >
      {children}
    </Box>
  );
}

Lighthouse ≥90 Checklist (build these in from the start)
AreaImplementationPerformanceLazy-load all route views with React.lazy + Suspense. Use loading="lazy" on all <img>. Serve images as WebP. No unused MUI icons (import individually, not barrel).AccessibilityAll interactive elements have aria-label. Color contrast ≥ 4.5:1. Navbar uses <nav> + aria-current="page". Skip-to-content link at top.SEOEach page sets <title> and <meta name="description"> via a useSEO hook or react-helmet-equivalent. Use semantic HTML (<main>, <section>, <article>, <h1>–<h3> hierarchy).Best PracticesHTTPS-ready. No console errors. Images have explicit width/height.
Lazy loading routes example:
tsxconst HomeView = React.lazy(() => import("../views/Home"));
// wrap in:
<Suspense fallback={<CircularProgress />}>
  <Outlet />
</Suspense>

Pages Summary
PageKey SectionsHomeHero (name, title, CTA buttons), Skills chips, featured Projects preview (3 cards)ExperienceVertical timeline of ExperienceCard componentsProjectsTag filter bar + responsive grid of ProjectCardContactSimple form (Name, Email, Message) + social links

Final Instructions for the Model

Scaffold the full folder structure above before writing any component.
Start with models/ → store/ → theme/ → router/ → controllers/ → components/ → views/ in that order.
Every placeholder string must be wrapped in a comment: {/* TODO: Replace with real content */}.
Do not install any packages. Use only what's listed in the tech stack.
Apply the glass theme consistently — every surface is a frosted panel, never a flat opaque box.
Import MUI icons individually (e.g. import GitHub from "@mui/icons-material/GitHub") to keep bundle size small.
Ensure all pages are accessible and semantically correct for Lighthouse.