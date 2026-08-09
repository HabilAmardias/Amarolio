import { Box, Typography, IconButton } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import GitHub from "@mui/icons-material/GitHub";
import LinkedIn from "@mui/icons-material/LinkedIn";
import Email from "@mui/icons-material/Email";

export default function Footer() {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  const socials = [
    {
      label: "GitHub",
      href: "https://github.com/HabilAmardias",
      icon: <GitHub fontSize="small" />,
    },
    {
      label: "LinkedIn",
      href: "https://www.linkedin.com/in/muhammad-habil-amardias/",
      icon: <LinkedIn fontSize="small" />,
    },
    {
      label: "Email",
      href: "mailto:habilamar@gmail.com",
      icon: <Email fontSize="small" />,
    },
  ];

  return (
    <Box
      component="footer"
      sx={{
        mt: "auto",
        py: 5,
        borderTop: `1px solid ${
          isLight ? "rgba(106,137,167,0.16)" : "rgba(189,221,252,0.08)"
        }`,
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: 2.5,
          px: 2,
        }}
      >
        <Box sx={{ display: "flex", gap: 1.5 }}>
          {socials.map((social) => (
            <IconButton
              key={social.label}
              href={social.href}
              aria-label={social.label}
              sx={{
                border: "1px solid",
                borderColor: "divider",
                color: "text.secondary",
                "&:hover": {
                  color: "primary.main",
                  borderColor: "primary.main",
                  transform: "translateY(-2px)",
                },
              }}
            >
              {social.icon}
            </IconButton>
          ))}
        </Box>
        <Typography
          variant="body2"
          color="text.secondary"
          sx={{ textAlign: "center" }}
        >
          © {new Date().getFullYear()} Amarolio. All rights reserved.{" "}
          {/* TODO: Replace with real content */}
        </Typography>
      </Box>
    </Box>
  );
}
