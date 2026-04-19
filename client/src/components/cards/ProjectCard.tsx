import { Card, CardContent, CardMedia, Typography, Box, Chip, IconButton } from "@mui/material";
import { GitHub, Launch } from "@mui/icons-material";
import type { Project } from "../../models/types";

interface ProjectCardProps {
  project: Project;
}

export default function ProjectCard({ project }: ProjectCardProps) {
  return (
    <Card sx={{ height: "100%", display: "flex", flexDirection: "column" }}>
      <CardMedia
        component="img"
        height="200"
        image={project.imageUrl}
        alt={project.title}
        loading="lazy"
        sx={{ objectFit: "cover" }}
      />
      <CardContent sx={{ flexGrow: 1, display: "flex", flexDirection: "column" }}>
        <Typography variant="h6" component="h3" gutterBottom>
          {project.title}
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ mb: 2, flexGrow: 1 }}>
          {project.description}
        </Typography>
        <Box sx={{ display: "flex", flexWrap: "wrap", gap: 0.5, mb: 2 }}>
          {project.tags.map((tag) => (
            <Chip key={tag} label={tag} size="small" variant="outlined" />
          ))}
        </Box>
        <Box sx={{ display: "flex", gap: 1 }}>
          {project.repoUrl && (
            <IconButton href={project.repoUrl} target="_blank" rel="noopener noreferrer" aria-label={`View ${project.title} repository`}>
              <GitHub />
            </IconButton>
          )}
          {project.liveUrl && (
            <IconButton href={project.liveUrl} target="_blank" rel="noopener noreferrer" aria-label={`View ${project.title} live`}>
              <Launch />
            </IconButton>
          )}
        </Box>
      </CardContent>
    </Card>
  );
}