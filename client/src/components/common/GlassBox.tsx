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
          ? "rgba(255,255,255,0.72)"
          : "rgba(35,43,58,0.55)",
        border: `1px solid ${
          isLight ? "rgba(61,107,212,0.14)" : "rgba(143,179,240,0.12)"
        }`,
        borderRadius: "30px",
        boxShadow: isLight
          ? "0 20px 50px -30px rgba(37,51,80,0.28)"
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
