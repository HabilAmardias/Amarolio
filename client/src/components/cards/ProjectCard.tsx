import {
  Card,
  CardContent,
  CardMedia,
  Typography,
  Box,
  Chip,
  IconButton,
} from "@mui/material";
import { GitHub, Launch } from "@mui/icons-material";
import type { Project } from "../../models/types";

interface ProjectCardProps {
  project: Project;
}

export default function ProjectCard({ project }: ProjectCardProps) {  return (
    <Card
      sx={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        overflow: "hidden",
      }}
    >
      <Box sx={{ position: "relative" }}>
        <CardMedia
          component="img"
          height="200"
          image={project.imageUrl}
          alt={project.title}
          loading="lazy"
          sx={{
            objectFit: "cover",
            display: "block",
            background: "rgba(106,137,167,0.12)",
          }}
        />
        {project.featured && (
          <Chip
            label="Featured"
            size="small"
            sx={{
              position: "absolute",
              top: 12,
              right: 12,
              bgcolor: "rgba(136,189,242,0.92)",
              color: "#0d1a22",
              fontWeight: 700,
              backdropFilter: "blur(6px)",
            }}
          />
        )}
      </Box>
      <CardContent
        sx={{
          flexGrow: 1,
          display: "flex",
          flexDirection: "column",
          p: 3,
        }}
      >
        <Typography
          variant="h6"
          component="h3"
          gutterBottom
          sx={{ fontWeight: 700, letterSpacing: "-0.01em" }}
        >
          {project.title}
        </Typography>
        <Typography
          variant="body2"
          color="text.secondary"
          sx={{ mb: 2, flexGrow: 1, lineHeight: 1.65 }}
        >
          {project.description}
        </Typography>
        <Box
          sx={{ display: "flex", flexWrap: "wrap", gap: 0.75, mb: 2.5 }}
        >
          {project.tags.map((tag) => (
            <Chip key={tag} label={tag} size="small" variant="outlined" />
          ))}
        </Box>
        <Box sx={{ display: "flex", gap: 1, alignItems: "center" }}>
          {project.repoUrl && (
            <IconButton
              href={project.repoUrl}
              target="_blank"
              rel="noopener noreferrer"
              aria-label={`View ${project.title} repository`}
              size="small"
              sx={{
                border: "1px solid",
                borderColor: "divider",
                color: "text.secondary",
                "&:hover": {
                  color: "primary.main",
                  borderColor: "primary.main",
                },
              }}
            >
              <GitHub />
            </IconButton>
          )}
          {project.liveUrl && project.liveUrl !== "#" && (
            <IconButton
              href={project.liveUrl}
              target="_blank"
              rel="noopener noreferrer"
              aria-label={`View ${project.title} live`}
              size="small"
              sx={{
                border: "1px solid",
                borderColor: "divider",
                color: "text.secondary",
                "&:hover": {
                  color: "primary.main",
                  borderColor: "primary.main",
                },
              }}
            >
              <Launch />
            </IconButton>
          )}
        </Box>
      </CardContent>
    </Card>
  );
}
