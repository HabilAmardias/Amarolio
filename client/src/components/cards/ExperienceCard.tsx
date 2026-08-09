import { Card, Typography, Box, Chip, Avatar } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import WorkOutlineRounded from "@mui/icons-material/WorkOutlineRounded";
import type { Experience } from "../../models/types";

interface ExperienceCardProps {
  experience: Experience;
}

export default function ExperienceCard({ experience }: ExperienceCardProps) {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  return (
    <Card
      className="animate-fade-up"
      sx={{
        display: "flex",
        gap: { xs: 2, sm: 3 },
        p: { xs: 2.5, sm: 3 },
        alignItems: "flex-start",
      }}
    >
      <Avatar
        sx={{
          width: { xs: 44, sm: 52 },
          height: { xs: 44, sm: 52 },
          borderRadius: "16px",
          flexShrink: 0,
          bgcolor: isLight
            ? "rgba(61,107,212,0.12)"
            : "rgba(143,179,240,0.16)",
          color: "primary.main",
          fontSize: { xs: "1.05rem", sm: "1.2rem" },
          fontWeight: 600,
          fontFamily: "'Poppins', sans-serif",
        }}
      >
        {experience.company.charAt(0).toUpperCase()}
      </Avatar>
      <Box sx={{ flexGrow: 1, minWidth: 0 }}>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            flexWrap: "wrap",
            gap: 1,
            alignItems: "baseline",
            mb: 0.5,
          }}
        >
          <Typography
            variant="h6"
            component="h3"
            sx={{
              fontWeight: 600,
              fontFamily: "'Poppins', sans-serif",
              letterSpacing: "-0.01em",
            }}
          >
            {experience.role}
          </Typography>
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
                ? "rgba(240,166,59,0.12)"
                : "rgba(240,166,59,0.1)",
            }}
          >
            <WorkOutlineRounded fontSize="small" sx={{ color: "warning.main" }} />
            <Typography variant="body2" color="text.secondary" sx={{ fontWeight: 600 }}>
              {experience.period}
            </Typography>
          </Box>
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
              <Chip
                key={tag}
                label={tag}
                size="small"
                variant="outlined"
                sx={{
                  "&:hover": {
                    bgcolor: isLight
                      ? "rgba(61,107,212,0.08)"
                      : "rgba(143,179,240,0.12)",
                  },
                }}
              />
            ))}
          </Box>
        )}
      </Box>
    </Card>
  );
}
