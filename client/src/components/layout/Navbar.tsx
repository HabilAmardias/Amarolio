import { useState } from "react";
import {
  AppBar,
  Toolbar,
  Typography,
  Button,
  Box,
  IconButton,
  Tooltip,
  Drawer,
  Divider,
  useMediaQuery,
} from "@mui/material";
import { useTheme } from "@mui/material/styles";
import { NavLink } from "react-router-dom";
import { useAtom } from "jotai";
import LightMode from "@mui/icons-material/LightMode";
import DarkMode from "@mui/icons-material/DarkMode";
import BrightnessAuto from "@mui/icons-material/BrightnessAuto";
import Menu from "@mui/icons-material/Menu";
import Close from "@mui/icons-material/Close";
import { themeModeAtom } from "../../store/atoms";
import type { ColorMode } from "../../theme/theme";
import type { ReactNode } from "react";

const navLinks = [
  { to: "/", label: "Home" },
  { to: "/experience", label: "Experience" },
  { to: "/projects", label: "Projects" },
];

const modeOrder: ColorMode[] = ["system", "light", "dark"];
const modeLabels: Record<ColorMode, string> = {
  system: "Theme: System",
  light: "Theme: Light",
  dark: "Theme: Dark",
};
const modeIcons: Record<ColorMode, ReactNode> = {
  system: <BrightnessAuto fontSize="small" />,
  light: <LightMode fontSize="small" />,
  dark: <DarkMode fontSize="small" />,
};

function Brand() {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  return (
    <Box
      component={NavLink}
      to="/"
      sx={{
        display: "flex",
        alignItems: "center",
        gap: 1.25,
        "&:hover .brand-badge": { transform: "rotate(-8deg) scale(1.06)" },
      }}
    >
      <Box
        className="brand-badge"
        sx={{
          width: 36,
          height: 36,
          borderRadius: "12px",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          position: "relative",
          color: isLight ? "#ffffff" : "#0E1A2E",
          fontWeight: 700,
          fontSize: "1.05rem",
          fontFamily: "'Poppins', sans-serif",
          background: isLight
            ? "linear-gradient(135deg, #3D6BD4, #8FB3F0)"
            : "linear-gradient(135deg, #C6DCFA, #8FB3F0)",
          boxShadow: isLight
            ? "0 8px 20px -8px rgba(61,107,212,0.6)"
            : "0 8px 20px -8px rgba(0,0,0,0.45)",
          transition: "transform 0.25s cubic-bezier(0.34, 1.56, 0.64, 1)",
          "&::after": {
            content: '""',
            position: "absolute",
            top: 2,
            right: 2,
            width: 8,
            height: 8,
            borderRadius: "50%",
            background: "#F0A63B",
            border: "2px solid",
            borderColor: isLight ? "#ffffff" : "#1B212E",
          },
        }}
      >
        A
      </Box>
      <Typography
        variant="h6"
        sx={{
          fontWeight: 600,
          fontFamily: "'Poppins', sans-serif",
          letterSpacing: "-0.02em",
          fontSize: "1.05rem",
        }}
      >
        Amarolio
      </Typography>
    </Box>
  );
}

function ThemeToggle() {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";
  const [colorMode, setColorMode] = useAtom(themeModeAtom);

  const cycleMode = () => {
    const index = modeOrder.indexOf(colorMode);
    setColorMode(modeOrder[(index + 1) % modeOrder.length]);
  };

  return (
    <Tooltip title={modeLabels[colorMode]}>
      <IconButton
        onClick={cycleMode}
        aria-label={`Theme: ${colorMode}`}
        size="small"
        sx={{
          color: "text.secondary",
          border: "1px solid",
          borderColor: "divider",
          borderRadius: 999,
          p: 0.9,
          transition: "all 0.2s ease",
          "&:hover": {
            color: "primary.main",
            borderColor: "primary.main",
            background: isLight
              ? "rgba(143,179,240,0.18)"
              : "rgba(143,179,240,0.14)",
          },
        }}
      >
        {modeIcons[colorMode]}
      </IconButton>
    </Tooltip>
  );
}

export default function Navbar() {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";
  const isMobile = useMediaQuery(theme.breakpoints.down("sm"));
  const [mobileOpen, setMobileOpen] = useState(false);

  const hoverBg = isLight ? "rgba(61,107,212,0.1)" : "rgba(143,179,240,0.14)";
  const activeBg = isLight
    ? "rgba(61,107,212,0.16)"
    : "rgba(143,179,240,0.18)";

  return (
    <AppBar position="sticky" color="transparent" elevation={0}>
      <Toolbar
        sx={{
          maxWidth: 1200,
          width: "100%",
          mx: "auto",
          px: { xs: 1.5, sm: 3 },
          minHeight: { xs: 60, sm: 72 },
          gap: 1,
        }}
      >
        <Box sx={{ flexGrow: 1, display: "flex", alignItems: "center" }}>
          <Brand />
        </Box>

        {!isMobile && (
          <nav aria-label="Main navigation">
            <Box sx={{ display: "flex", gap: 0.5, alignItems: "center" }}>
              {navLinks.map((link) => (
                <Button
                  key={link.to}
                  component={NavLink}
                  to={link.to}
                  end={link.to === "/"}
                  sx={{
                    color: "text.secondary",
                    px: 1.75,
                    py: 0.75,
                    borderRadius: 999,
                    fontWeight: 500,
                    fontSize: "0.9rem",
                    whiteSpace: "nowrap",
                    position: "relative",
                    transition: "all 0.2s ease",
                    "&:hover": {
                      color: "text.primary",
                      background: hoverBg,
                      transform: "translateY(-1px)",
                    },
                    "&[aria-current='page']": {
                      color: "primary.main",
                      fontWeight: 700,
                      background: activeBg,
                      "&::after": {
                        content: '""',
                        position: "absolute",
                        top: "50%",
                        right: 4,
                        transform: "translateY(-50%)",
                        width: 6,
                        height: 6,
                        borderRadius: "50%",
                        background: "#F0A63B",
                      },
                    },
                  }}
                >
                  {link.label}
                </Button>
              ))}
            </Box>
          </nav>
        )}

        <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
          <ThemeToggle />
          {isMobile && (
            <IconButton
              onClick={() => setMobileOpen(true)}
              aria-label="Open navigation menu"
              size="small"
              sx={{
                color: "text.secondary",
                border: "1px solid",
                borderColor: "divider",
                borderRadius: 999,
                p: 0.9,
                "&:hover": {
                  color: "primary.main",
                  borderColor: "primary.main",
                },
              }}
            >
              <Menu />
            </IconButton>
          )}
        </Box>
      </Toolbar>

      <Drawer anchor="right" open={mobileOpen} onClose={() => setMobileOpen(false)}>
        <Box
          role="navigation"
          aria-label="Mobile navigation"
          sx={{
            width: 280,
            height: "100%",
            display: "flex",
            flexDirection: "column",
            px: 2,
            py: 2,
          }}
        >
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              mb: 1,
            }}
          >
            <Brand />
            <IconButton
              onClick={() => setMobileOpen(false)}
              aria-label="Close navigation menu"
              size="small"
              sx={{
                color: "text.secondary",
                "&:hover": { color: "primary.main" },
              }}
            >
              <Close />
            </IconButton>
          </Box>

          <Divider sx={{ my: 1.5 }} />

          <Box sx={{ display: "flex", flexDirection: "column", gap: 0.5 }}>
            {navLinks.map((link) => (
              <Button
                key={link.to}
                component={NavLink}
                to={link.to}
                end={link.to === "/"}
                onClick={() => setMobileOpen(false)}
                fullWidth
                sx={{
                  justifyContent: "flex-start",
                  color: "text.secondary",
                  px: 2,
                  py: 1.25,
                  borderRadius: 999,
                  fontWeight: 500,
                  fontSize: "1rem",
                  "&:hover": {
                    color: "text.primary",
                    background: hoverBg,
                  },
                  "&[aria-current='page']": {
                    color: "primary.main",
                    fontWeight: 700,
                    background: activeBg,
                  },
                }}
              >
                {link.label}
              </Button>
            ))}
          </Box>
        </Box>
      </Drawer>
    </AppBar>
  );
}
