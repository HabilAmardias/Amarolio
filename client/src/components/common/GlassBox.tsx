import { Box } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import type { BoxProps } from "@mui/material";

export default function GlassBox({ children, sx, ...props }: BoxProps) {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  return (
    <Box
      sx={{
        backdropFilter: "blur(18px) saturate(150%)",
        background: isLight
          ? "rgba(255,255,255,0.75)"
          : "rgba(30,44,55,0.55)",
        border: `1px solid ${
          isLight ? "rgba(106,137,167,0.16)" : "rgba(189,221,252,0.1)"
        }`,
        borderRadius: "28px",
        boxShadow: isLight
          ? "0 20px 50px -30px rgba(56,73,89,0.3)"
          : "0 20px 50px -24px rgba(0,0,0,0.5)",
        p: 4,
        ...sx,
      }}
      {...props}
    >
      {children}
    </Box>
  );
}
