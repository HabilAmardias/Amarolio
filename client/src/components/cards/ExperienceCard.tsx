import { Card, CardContent, Typography, Box, Chip } from "@mui/material";
import type { Experience } from "../../models/types";

interface ExperienceCardProps {
  experience: Experience;
}

export default function ExperienceCard({ experience }: ExperienceCardProps) {
  return (
    <Card sx={{ position: "relative", pl: 3, borderLeft: "3px solid", borderColor: "primary.main" }}>
      <CardContent>
        <Typography variant="h6" component="h3" gutterBottom>
          {experience.role}
        </Typography>
        <Typography variant="subtitle1" color="primary" gutterBottom>
          {experience.company}
        </Typography>
        <Box sx={{ display: "flex", justifyContent: "space-between", mb: 2, flexWrap: "wrap", gap: 1 }}>
          <Typography variant="body2" color="text.secondary">
            {experience.period}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            {experience.location}
          </Typography>
        </Box>
        <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
          {experience.description}
        </Typography>
        <Box sx={{ display: "flex", flexWrap: "wrap", gap: 0.5 }}>
          {experience.tags.map((tag) => (
            <Chip key={tag} label={tag} size="small" variant="outlined" />
          ))}
        </Box>
      </CardContent>
    </Card>
  );
}