import { Typography, Box } from "@mui/material";
import { useTheme } from "@mui/material/styles";

interface SectionHeaderProps {
  eyebrow?: string;
  title: string;
  subtitle?: string;
}

export default function SectionHeader({
  eyebrow,
  title,
  subtitle,
}: SectionHeaderProps) {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  return (
    <Box
      className="animate-fade-up"
      sx={{
        mb: 5,
        textAlign: { xs: "center", sm: "left" },
        maxWidth: 720,
        mx: { xs: "auto", sm: 0 },
      }}
    >
      {eyebrow && (
        <Box
          sx={{
            display: "inline-flex",
            alignItems: "center",
            gap: 0.75,
            px: 1.25,
            py: 0.5,
            borderRadius: 999,
            border: "1px solid",
            borderColor: "divider",
            background: isLight
              ? "rgba(143,179,240,0.16)"
              : "rgba(143,179,240,0.12)",
            color: "primary.main",
            fontSize: "0.78rem",
            fontWeight: 600,
            textTransform: "uppercase",
            letterSpacing: "0.1em",
            mb: 1.5,
          }}
        >
          <Box
            sx={{
              width: 6,
              height: 6,
              borderRadius: "50%",
              bgcolor: "warning.main",
            }}
          />
          {eyebrow}
        </Box>
      )}
      <Typography
        variant="h2"
        component="h1"
        gutterBottom
        sx={{ fontSize: { xs: "1.7rem", sm: "2rem" } }}
      >
        {title}
      </Typography>
      {subtitle && (
        <Typography
          variant="body1"
          color="text.secondary"
          sx={{ lineHeight: 1.7 }}
        >
          {subtitle}
        </Typography>
      )}
    </Box>
  );
}
