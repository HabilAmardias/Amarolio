import { createTheme } from "@mui/material/styles";
import type { Theme } from "@mui/material/styles";
import { useEffect, useState } from "react";

export type ColorMode = "system" | "light" | "dark";

export const brand = {
  blue: "#3D6BD4",
  sky: "#8FB3F0",
  mist: "#DCE8FB",
  navy: "#253350",
  ink: "#222A3B",
  slate: "#5C6B84",
  amber: "#F0A63B",
  cream: "#F7F4EE",
};

const typography = {
  fontFamily: "'Open Sans', 'Poppins', 'Inter', sans-serif",
  h1: {
    fontFamily: "'Poppins', 'Open Sans', sans-serif",
    fontWeight: 700,
    letterSpacing: "-0.02em",
    lineHeight: 1.12,
  },
  h2: {
    fontFamily: "'Poppins', 'Open Sans', sans-serif",
    fontWeight: 700,
    letterSpacing: "-0.015em",
    lineHeight: 1.18,
  },
  h3: {
    fontFamily: "'Poppins', 'Open Sans', sans-serif",
    fontWeight: 600,
    letterSpacing: "-0.01em",
    lineHeight: 1.28,
  },
  h4: {
    fontFamily: "'Poppins', 'Open Sans', sans-serif",
    fontWeight: 600,
    letterSpacing: "-0.01em",
  },
  h5: {
    fontFamily: "'Poppins', 'Open Sans', sans-serif",
    fontWeight: 600,
  },
  h6: {
    fontFamily: "'Poppins', 'Open Sans', sans-serif",
    fontWeight: 600,
  },
  button: { textTransform: "none", fontWeight: 600 },
  body1: { lineHeight: 1.7 },
  body2: { lineHeight: 1.65 },
};

function buildTheme(mode: "light" | "dark"): Theme {
  const isLight = mode === "light";
  return createTheme({
    palette: {
      mode,
      ...(isLight
        ? {
            background: { default: brand.cream, paper: "#ffffff" },
            primary: {
              main: brand.blue,
              light: brand.sky,
              dark: brand.navy,
              contrastText: "#ffffff",
            },
            secondary: {
              main: brand.sky,
              light: brand.mist,
              contrastText: brand.navy,
            },
            warning: {
              main: brand.amber,
              contrastText: "#4A3210",
            },
            text: { primary: brand.ink, secondary: brand.slate },
            divider: "rgba(61, 107, 212, 0.16)",
          }
        : {
            background: { default: "#1B212E", paper: "#232B3A" },
            primary: {
              main: brand.sky,
              light: "#C6DCFA",
              dark: brand.blue,
              contrastText: "#0E1A2E",
            },
            secondary: {
              main: "#6F8FC8",
              light: brand.sky,
              contrastText: "#0E1A2E",
            },
            warning: {
              main: "#F5B45C",
              contrastText: "#241B08",
            },
            text: { primary: "#F2F5FA", secondary: "#A9B7CD" },
            divider: "rgba(143, 179, 240, 0.16)",
          }),
    },
    typography,
    shape: { borderRadius: 14 },
    components: {
      MuiCssBaseline: {
        styleOverrides: {
          body: {
            fontFamily: "'Open Sans', 'Poppins', 'Inter', sans-serif",
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
            background: isLight ? "#ffffff" : "#232B3A",
            border: `1px solid ${
              isLight ? "rgba(61,107,212,0.14)" : "rgba(143,179,240,0.14)"
            }`,
            borderRadius: 22,
            transition:
              "transform 0.28s cubic-bezier(0.22, 1, 0.36, 1), box-shadow 0.28s cubic-bezier(0.22, 1, 0.36, 1), border-color 0.28s ease",
            "&:hover": {
              transform: "translateY(-5px)",
              boxShadow: isLight
                ? "0 20px 44px -18px rgba(37,51,80,0.28)"
                : "0 20px 48px -14px rgba(0,0,0,0.55)",
              borderColor: isLight
                ? "rgba(61,107,212,0.35)"
                : "rgba(143,179,240,0.35)",
            },
          },
        },
      },
      MuiAppBar: {
        styleOverrides: {
          root: {
            background: isLight
              ? "rgba(247,244,238,0.72)"
              : "rgba(27,33,46,0.72)",
            backdropFilter: "blur(18px) saturate(150%)",
            borderBottom: `1px solid ${
              isLight ? "rgba(61,107,212,0.1)" : "rgba(143,179,240,0.1)"
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
            transition:
              "transform 0.2s cubic-bezier(0.22, 1, 0.36, 1), box-shadow 0.2s ease, background-color 0.2s ease",
            "&:active": {
              transform: "translateY(1px) scale(0.98)",
            },
          },
          contained: {
            boxShadow: isLight
              ? "0 10px 24px -10px rgba(61,107,212,0.55)"
              : "0 8px 24px -8px rgba(0,0,0,0.5)",
            "&:hover": {
              boxShadow: isLight
                ? "0 14px 30px -12px rgba(61,107,212,0.7)"
                : "0 10px 28px -8px rgba(0,0,0,0.6)",
            },
          },
        },
      },
      MuiChip: {
        styleOverrides: {
          root: {
            fontWeight: 600,
            borderRadius: 999,
          },
          outlined: {
            borderColor: isLight
              ? "rgba(61,107,212,0.35)"
              : "rgba(143,179,240,0.3)",
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
