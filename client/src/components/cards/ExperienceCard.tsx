import { Card, Typography, Box, Chip, Avatar } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import type { Experience } from "../../models/types";

interface ExperienceCardProps {
  experience: Experience;
}

export default function ExperienceCard({ experience }: ExperienceCardProps) {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  const initials = experience.company
    .split(" ")
    .filter(Boolean)
    .slice(0, 2)
    .map((word) => word.charAt(0).toUpperCase())
    .join("");

  return (
    <Card
      sx={{
        display: "flex",
        gap: { xs: 2, sm: 3 },
        p: { xs: 2.5, sm: 3 },
        position: "relative",
        overflow: "hidden",
        "&::before": {
          content: '""',
          position: "absolute",
          top: 0,
          bottom: 0,
          left: 0,
          width: 4,
          background: isLight
            ? "linear-gradient(180deg, #88bdf2, #6a89a7)"
            : "linear-gradient(180deg, #bdddfc, #88bdf2)",
        },
      }}
    >
      <Box
        sx={{
          display: { xs: "none", sm: "flex" },
          flexDirection: "column",
          alignItems: "center",
        }}
      >
      </Box>
      <Box sx={{ flexGrow: 1, minWidth: 0 }}>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            flexWrap: "wrap",
            gap: 1,
            alignItems: "baseline",
            mb: 0.75,
          }}
        >
          <Typography
            variant="h6"
            component="h3"
            sx={{ fontWeight: 700, letterSpacing: "-0.01em" }}
          >
            {experience.role}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {experience.period}
          </Typography>
        </Box>
        <Box sx={{ display: "flex", flexWrap: "wrap", gap: 1.5, mb: 1.25 }}>
          <Typography
            variant="subtitle1"
            color="primary"
            sx={{ fontWeight: 600 }}
          >
            {experience.company}
          </Typography>
          <Typography
            variant="subtitle1"
            color="text.secondary"
            sx={{ fontWeight: 400 }}
          >
            {experience.location}
          </Typography>
        </Box>
        <Typography
          variant="body2"
          color="text.secondary"
          sx={{ mb: 2, lineHeight: 1.7 }}
        >
          {experience.description}
        </Typography>
        {experience.tags.length > 0 && (
          <Box sx={{ display: "flex", flexWrap: "wrap", gap: 0.75 }}>
            {experience.tags.map((tag) => (
              <Chip key={tag} label={tag} size="small" variant="outlined" />
            ))}
          </Box>
        )}
      </Box>
    </Card>
  );
}
