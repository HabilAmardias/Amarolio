import { Box, Typography, IconButton } from "@mui/material";
import GitHub from "@mui/icons-material/GitHub";
import LinkedIn from "@mui/icons-material/LinkedIn";
import Email from "@mui/icons-material/Email";

export default function Footer() {
  return (
    <Box
      component="footer"
      sx={{
        py: 4,
        textAlign: "center",
        borderTop: "1px solid rgba(255,255,255,0.06)",
        mt: "auto",
      }}
    >
      <Box sx={{ display: "flex", justifyContent: "center", gap: 2, mb: 2 }}>
        <IconButton href="https://github.com/HabilAmardias" aria-label="GitHub" color="inherit">
          <GitHub />
        </IconButton>
        <IconButton href="https://www.linkedin.com/in/muhammad-habil-amardias/" aria-label="LinkedIn" color="inherit">
          <LinkedIn />
        </IconButton>
        <IconButton href="mailto:habilamar@gmail.com" aria-label="Email" color="inherit">
          <Email />
        </IconButton>
      </Box>
      <Typography variant="body2" color="text.secondary">
        © {new Date().getFullYear()} Amarolio. All rights reserved. {/* TODO: Replace with real content */}
      </Typography>
    </Box>
  );
}