import { createTheme } from "@mui/material/styles";
import type { Theme } from "@mui/material/styles";
import { useEffect, useState } from "react";

export type ColorMode = "system" | "light" | "dark";

export const brand = {
  steel: "#6a89a7",
  pale: "#bdddfc",
  sky: "#88bdf2",
  navy: "#384959",
};

const typography = {
  fontFamily: "'Inter', 'Plus Jakarta Sans', 'Sora', 'DM Sans', sans-serif",
  h1: {
    fontFamily: "'Plus Jakarta Sans', 'Inter', sans-serif",
    fontWeight: 800,
    letterSpacing: "-0.02em",
    lineHeight: 1.12,
  },
  h2: {
    fontFamily: "'Plus Jakarta Sans', 'Inter', sans-serif",
    fontWeight: 800,
    letterSpacing: "-0.015em",
    lineHeight: 1.15,
  },
  h3: {
    fontFamily: "'Plus Jakarta Sans', 'Inter', sans-serif",
    fontWeight: 700,
    letterSpacing: "-0.01em",
    lineHeight: 1.25,
  },
  h4: {
    fontFamily: "'Plus Jakarta Sans', 'Inter', sans-serif",
    fontWeight: 700,
    letterSpacing: "-0.01em",
  },
  h5: {
    fontFamily: "'Plus Jakarta Sans', 'Inter', sans-serif",
    fontWeight: 700,
  },
  h6: {
    fontFamily: "'Plus Jakarta Sans', 'Inter', sans-serif",
    fontWeight: 700,
  },
  button: { textTransform: "none", fontWeight: 600 },
};

function buildTheme(mode: "light" | "dark"): Theme {
  const isLight = mode === "light";
  return createTheme({
    palette: {
      mode,
      ...(isLight
        ? {
            background: { default: "#f4f7fb", paper: "#ffffff" },
            primary: {
              main: brand.steel,
              light: brand.sky,
              dark: brand.navy,
              contrastText: "#ffffff",
            },
            secondary: {
              main: brand.sky,
              light: brand.pale,
              contrastText: "#18303f",
            },
            text: { primary: "#384959", secondary: "#4e6b87" },
            divider: "rgba(106, 137, 167, 0.22)",
          }
        : {
            background: { default: "#152129", paper: "#1e2c37" },
            primary: {
              main: brand.sky,
              light: brand.pale,
              dark: brand.steel,
              contrastText: "#0d1a22",
            },
            secondary: {
              main: brand.steel,
              light: brand.sky,
              contrastText: "#f0f7fe",
            },
            text: { primary: "#eaf3fb", secondary: "#a9c0d5" },
            divider: "rgba(189, 221, 252, 0.14)",
          }),
    },
    typography,
    shape: { borderRadius: 12 },
    components: {
      MuiCssBaseline: {
        styleOverrides: {
          body: {
            fontFamily: "'Inter', 'Plus Jakarta Sans', 'Sora', 'DM Sans', sans-serif",
          },
        },
      },
      MuiPaper: {
        styleOverrides: {
          root: {
            backgroundImage: "none",
          },
        },
      },
      MuiCard: {
        styleOverrides: {
          root: {
            backgroundImage: "none",
            background: isLight ? "#ffffff" : "#1e2c37",
            border: `1px solid ${
              isLight ? "rgba(106,137,167,0.18)" : "rgba(189,221,252,0.12)"
            }`,
            borderRadius: 20,
            transition:
              "transform 0.25s ease, box-shadow 0.25s ease, border-color 0.25s ease",
            "&:hover": {
              transform: "translateY(-4px)",
              boxShadow: isLight
                ? "0 18px 40px -18px rgba(56,73,89,0.3)"
                : "0 18px 48px -12px rgba(0,0,0,0.55)",
              borderColor: isLight
                ? "rgba(136,189,242,0.5)"
                : "rgba(189,221,252,0.3)",
            },
          },
        },
      },
      MuiAppBar: {
        styleOverrides: {
          root: {
            background: isLight
              ? "rgba(244,247,251,0.75)"
              : "rgba(21,33,41,0.72)",
            backdropFilter: "blur(16px) saturate(160%)",
            borderBottom: `1px solid ${
              isLight ? "rgba(106,137,167,0.14)" : "rgba(189,221,252,0.1)"
            }`,
            boxShadow: "none",
          },
        },
      },
      MuiButton: {
        styleOverrides: {
          root: {
            borderRadius: 999,
            textTransform: "none",
            fontWeight: 600,
            letterSpacing: "0.01em",
          },
          contained: {
            boxShadow: isLight
              ? "0 10px 24px -10px rgba(56,73,89,0.55)"
              : "0 8px 24px -8px rgba(0,0,0,0.5)",
            "&:hover": {
              boxShadow: isLight
                ? "0 12px 28px -10px rgba(56,73,89,0.7)"
                : "0 10px 28px -8px rgba(0,0,0,0.6)",
            },
          },
        },
      },
      MuiChip: {
        styleOverrides: {
          root: {
            fontWeight: 500,
            borderRadius: 999,
          },
          outlined: {
            borderColor: isLight
              ? "rgba(106,137,167,0.4)"
              : "rgba(189,221,252,0.25)",
          },
        },
      },
    },
  });
}

export const lightTheme = buildTheme("light");
export const darkTheme = buildTheme("dark");

export function useSystemTheme(): "light" | "dark" {
  const [mode, setMode] = useState<"light" | "dark">(() =>
    typeof window !== "undefined" &&
    window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light"
  );

  useEffect(() => {
    const mql = window.matchMedia("(prefers-color-scheme: dark)");
    const onChange = (event: MediaQueryListEvent) =>
      setMode(event.matches ? "dark" : "light");
    mql.addEventListener("change", onChange);
    return () => mql.removeEventListener("change", onChange);
  }, []);

  return mode;
}
