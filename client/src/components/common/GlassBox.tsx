import { Box } from "@mui/material";
import type { BoxProps } from "@mui/material";

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