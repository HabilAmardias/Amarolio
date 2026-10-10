import { Box, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";

interface BrandProps {
  subtitle?: string;
}

export function Brand({ subtitle }: BrandProps) {
  return (
    <Box sx={{ display: "flex", alignItems: "center", gap: 1.25 }}>
      <Box
        aria-hidden
        sx={{
          width: 40,
          height: 40,
          borderRadius: 2.5,
          display: "grid",
          placeItems: "center",
          background: "linear-gradient(160deg, #8FB3F0 0%, #3D6BD4 100%)",
          color: "#ffffff",
          fontWeight: 800,
          fontSize: "1.15rem",
          fontFamily: "'Poppins', 'Open Sans', sans-serif",
          boxShadow: "0 10px 24px -12px rgba(61,107,212,0.6)",
        }}
      >
        A
      </Box>
      <Box sx={{ lineHeight: 1.1 }}>
        <Typography
          component={RouterLink}
          to="/"
          sx={{
            fontFamily: "'Poppins', 'Open Sans', sans-serif",
            fontWeight: 700,
            fontSize: "1.15rem",
            color: "text.primary",
            letterSpacing: "-0.01em",
          }}
        >
          Amarolio
        </Typography>
        {subtitle && (
          <Typography
            variant="caption"
            sx={{ display: "block", color: "text.secondary" }}
          >
            {subtitle}
          </Typography>
        )}
      </Box>
    </Box>
  );
}
