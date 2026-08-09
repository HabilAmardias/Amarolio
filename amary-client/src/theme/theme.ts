import { createTheme } from "@mui/material/styles";

// ---------------------------------------------------------------------------
// Amary Autumn Liquid Glass Design Tokens
// Fluid translucent surfaces: high-saturation backdrop blur, iridescent
// autumn gradients, smooth 400ms transitions, rounded 20px. Warm palette kept.
// ---------------------------------------------------------------------------
export const liquid = {
  rust: "#B45309", // primary — burnt amber
  rustDark: "#8A3E06",
  rustLight: "#E8820F",
  saddle: "#92400E", // secondary — warm brown
  saddleDark: "#5A2C0C",
  amber: "#F59E0B", // reward gold accent
  peach: "#F9C08A", // light accent
  cream: "#FFF8EF", // base canvas
  ink: "#3B2417", // primary text
  inkSoft: "#7A5C44", // secondary text
  leaf: "#3E7C59", // success
  rose: "#C0392B", // error
  // Liquid glass surfaces
  glassSurface: "linear-gradient(145deg, rgba(255, 255, 255, 0.6) 0%, rgba(255, 244, 230, 0.35) 100%)",
  glassBorder: "rgba(255, 255, 255, 0.65)",
  glassShadow:
    "0 8px 32px rgba(120, 53, 15, 0.14), inset 0 1px 0 rgba(255, 255, 255, 0.7)",
  glassBlur: "blur(24px) saturate(160%)",
} as const;

// Vibrant warm autumn canvas that the liquid glass diffuses and saturates.
export const liquidBackground =
  "radial-gradient(1100px 560px at 8% -8%, #FFE8C2 0%, transparent 55%)," +
  "radial-gradient(900px 560px at 100% 4%, #F9C08A 0%, transparent 52%)," +
  "radial-gradient(1000px 640px at 55% 112%, #F3A45B 0%, transparent 55%)," +
  "linear-gradient(180deg, #FFF8EF 0%, #F8E7D2 100%)";

export const theme = createTheme({
  palette: {
    mode: "light",
    primary: {
      main: liquid.rust,
      dark: liquid.rustDark,
      light: liquid.rustLight,
      contrastText: "#FFFFFF",
    },
    secondary: {
      main: liquid.saddle,
      dark: liquid.saddleDark,
      light: liquid.amber,
      contrastText: "#FFFFFF",
    },
    background: {
      default: liquid.cream,
      paper: "#FFFFFF",
    },
    text: {
      primary: liquid.ink,
      secondary: liquid.inkSoft,
    },
    divider: "rgba(180, 83, 9, 0.16)",
    error: { main: liquid.rose },
    success: { main: liquid.leaf },
    warning: { main: liquid.amber },
  },
  typography: {
    fontFamily: "'Inter', 'Helvetica Neue', Arial, sans-serif",
    h1: {
      fontFamily: "'Merriweather', 'Georgia', serif",
      fontWeight: 700,
      letterSpacing: "-0.5px",
    },
    h2: {
      fontFamily: "'Merriweather', 'Georgia', serif",
      fontWeight: 700,
      letterSpacing: "-0.5px",
    },
    h3: {
      fontFamily: "'Merriweather', 'Georgia', serif",
      fontWeight: 700,
      letterSpacing: "-0.3px",
    },
    h4: {
      fontFamily: "'Merriweather', 'Georgia', serif",
      fontWeight: 700,
      letterSpacing: "-0.2px",
    },
    h5: {
      fontFamily: "'Merriweather', 'Georgia', serif",
      fontWeight: 600,
    },
    h6: {
      fontFamily: "'Merriweather', 'Georgia', serif",
      fontWeight: 600,
    },
    button: {
      textTransform: "none",
      fontWeight: 600,
      letterSpacing: "0.2px",
    },
  },
  shape: { borderRadius: 20 },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        body: {
          background: liquidBackground,
          minHeight: "100vh",
          backgroundAttachment: "fixed",
        },
        "::selection": {
          backgroundColor: "rgba(180, 83, 9, 0.2)",
          color: "#3B2417",
        },
      },
    },
    MuiButton: {
      defaultProps: { disableElevation: true },
      styleOverrides: {
        root: {
          borderRadius: 14,
          fontWeight: 600,
          transition: "all 0.4s cubic-bezier(0.4, 0, 0.2, 1)",
        },
        contained: {
          background: `linear-gradient(160deg, ${liquid.rustLight} 0%, ${liquid.rust} 60%, ${liquid.rustDark} 100%)`,
          color: "#FFFFFF",
          border: "1px solid rgba(255, 255, 255, 0.35)",
          boxShadow: `0 8px 24px rgba(180, 83, 9, 0.32), inset 0 1px 0 rgba(255, 255, 255, 0.4)`,
          "&:hover": {
            boxShadow: `0 12px 32px rgba(180, 83, 9, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.45)`,
            transform: "translateY(-1px)",
          },
          "&:active": {
            transform: "translateY(1px)",
            boxShadow: `0 4px 12px rgba(180, 83, 9, 0.25)`,
          },
          "&.Mui-disabled": {
            background: "rgba(255, 255, 255, 0.6)",
            color: "rgba(62, 43, 29, 0.4)",
            boxShadow: "none",
          },
        },
        outlined: {
          backgroundColor: "rgba(255, 255, 255, 0.4)",
          backdropFilter: "blur(12px) saturate(140%)",
          WebkitBackdropFilter: "blur(12px) saturate(140%)",
          borderColor: "rgba(255, 255, 255, 0.8)",
          color: liquid.rust,
          boxShadow: "inset 0 1px 0 rgba(255, 255, 255, 0.6)",
          "&:hover": {
            backgroundColor: "rgba(255, 255, 255, 0.6)",
            borderColor: liquid.rust,
          },
        },
        text: {
          color: liquid.rust,
          "&:hover": {
            backgroundColor: "rgba(180, 83, 9, 0.1)",
          },
        },
      },
    },
    MuiAppBar: {
      defaultProps: { color: "inherit", elevation: 0 },
      styleOverrides: {
        root: {
          backgroundColor: "rgba(255, 250, 242, 0.5)",
          backdropFilter: "blur(24px) saturate(160%)",
          WebkitBackdropFilter: "blur(24px) saturate(160%)",
          borderBottom: `1px solid ${liquid.glassBorder}`,
          boxShadow: "inset 0 1px 0 rgba(255, 255, 255, 0.6)",
        },
      },
    },
    MuiTextField: {
      defaultProps: { size: "small", variant: "outlined" },
      styleOverrides: {
        root: {
          "& .MuiOutlinedInput-root": {
            background: "linear-gradient(145deg, rgba(255,255,255,0.75) 0%, rgba(255, 244, 230, 0.45) 100%)",
            backdropFilter: liquid.glassBlur,
            WebkitBackdropFilter: liquid.glassBlur,
            borderRadius: 14,
            boxShadow: "inset 0 1px 0 rgba(255, 255, 255, 0.7)",
            transition: "all 0.4s cubic-bezier(0.4, 0, 0.2, 1)",
            "& .MuiOutlinedInput-notchedOutline": {
              borderColor: "rgba(255, 255, 255, 0.7)",
            },
            "&:hover .MuiOutlinedInput-notchedOutline": {
              borderColor: liquid.rustLight,
            },
            "&.Mui-focused .MuiOutlinedInput-notchedOutline": {
              borderColor: liquid.rust,
              borderWidth: 1.5,
            },
          },
          "& .MuiInputLabel-root.Mui-focused": {
            color: liquid.rust,
          },
        },
      },
    },
    MuiCard: {
      styleOverrides: {
        root: {
          background: liquid.glassSurface,
          backdropFilter: liquid.glassBlur,
          WebkitBackdropFilter: liquid.glassBlur,
          border: `1px solid ${liquid.glassBorder}`,
          borderRadius: 22,
          boxShadow: liquid.glassShadow,
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundColor: liquid.cream,
          borderRadius: 22,
        },
        outlined: {
          background: liquid.glassSurface,
          border: `1px solid ${liquid.glassBorder}`,
          boxShadow: liquid.glassShadow,
        },
        elevation1: {
          background: liquid.glassSurface,
          boxShadow: liquid.glassShadow,
        },
        elevation2: {
          background: liquid.glassSurface,
          boxShadow: liquid.glassShadow,
        },
      },
    },
    MuiTable: {
      styleOverrides: {
        root: {
          minWidth: 640,
        },
      },
    },
    MuiTableCell: {
      styleOverrides: {
        head: {
          fontWeight: 700,
          fontSize: "0.78rem",
          letterSpacing: "0.4px",
          textTransform: "uppercase",
          color: liquid.rust,
          backgroundColor: "rgba(255, 248, 239, 0.5)",
          borderBottom: `1px solid ${liquid.glassBorder}`,
        },
        root: {
          borderBottom: "1px solid rgba(180, 83, 9, 0.12)",
        },
      },
    },
    MuiTableRow: {
      styleOverrides: {
        root: {
          "&:last-child .MuiTableCell-root": {
            borderBottom: "none",
          },
          "&:hover": {
            backgroundColor: "rgba(255, 255, 255, 0.45)",
          },
        },
      },
    },
    MuiAlert: {
      styleOverrides: {
        root: {
          borderRadius: 14,
          fontWeight: 600,
          backdropFilter: "blur(14px) saturate(140%)",
          WebkitBackdropFilter: "blur(14px) saturate(140%)",
        },
        standardError: {
          backgroundColor: "rgba(192, 57, 43, 0.12)",
          color: "#8B1E1A",
          border: "1px solid rgba(192, 57, 43, 0.25)",
        },
      },
    },
    MuiTabs: {
      styleOverrides: {
        indicator: {
          backgroundColor: liquid.rust,
          height: 3,
          borderRadius: 3,
        },
      },
    },
    MuiTab: {
      styleOverrides: {
        root: {
          textTransform: "none",
          fontWeight: 600,
          color: liquid.inkSoft,
          borderRadius: 10,
          "&.Mui-selected": {
            color: liquid.rust,
          },
        },
      },
    },
    MuiTooltip: {
      styleOverrides: {
        tooltip: {
          backgroundColor: "rgba(90, 44, 12, 0.85)",
          backdropFilter: "blur(12px) saturate(140%)",
          WebkitBackdropFilter: "blur(12px) saturate(140%)",
          fontSize: "0.78rem",
          borderRadius: 10,
          border: "1px solid rgba(255, 255, 255, 0.2)",
          boxShadow: "0 8px 20px rgba(0, 0, 0, 0.2)",
        },
        arrow: {
          color: "rgba(90, 44, 12, 0.85)",
        },
      },
    },
    MuiSnackbar: {
      styleOverrides: {
        root: {
          "& .MuiPaper-root": {
            backgroundColor: "rgba(90, 44, 12, 0.85)",
            backdropFilter: "blur(14px) saturate(140%)",
            WebkitBackdropFilter: "blur(14px) saturate(140%)",
            color: "#FFFFFF",
            borderRadius: 14,
            border: "1px solid rgba(255, 255, 255, 0.2)",
          },
        },
      },
    },
    MuiIconButton: {
      styleOverrides: {
        root: {
          borderRadius: 12,
          color: liquid.rust,
          transition: "all 0.3s cubic-bezier(0.4, 0, 0.2, 1)",
          "&:hover": {
            backgroundColor: "rgba(255, 255, 255, 0.5)",
          },
        },
      },
    },
    MuiCheckbox: {
      styleOverrides: {
        root: {
          borderRadius: 6,
          color: liquid.inkSoft,
          "&.Mui-checked": {
            color: liquid.rust,
          },
        },
      },
    },
    MuiTablePagination: {
      styleOverrides: {
        root: {
          color: liquid.inkSoft,
        },
        selectIcon: {
          color: liquid.rust,
        },
      },
    },
  },
});
