import {
  Card,
  CardContent,
  CardMedia,
  Typography,
  Box,
  Chip,
  IconButton,
} from "@mui/material";
import { useTheme } from "@mui/material/styles";
import { GitHub, Launch } from "@mui/icons-material";
import type { Project } from "../../models/types";

interface ProjectCardProps {
  project: Project;
}

export default function ProjectCard({ project }: ProjectCardProps) {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";

  return (
    <Card
      className="animate-fade-up"
      sx={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        overflow: "hidden",
        "& .card-media": {
          transition: "transform 0.5s cubic-bezier(0.22,1,0.36,1)",
        },
        "&:hover .card-media": { transform: "scale(1.07)" },
      }}
    >
      <Box sx={{ position: "relative", overflow: "hidden" }}>
        <CardMedia
          className="card-media"
          component="img"
          height="200"
          image={project.imageUrl}
          alt={project.title}
          loading="lazy"
          sx={{
            objectFit: "cover",
            display: "block",
            background: "rgba(61,107,212,0.1)",
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
              bgcolor: "#F0A63B",
              color: "#4A3210",
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
          sx={{
            fontWeight: 600,
            fontFamily: "'Poppins', sans-serif",
            letterSpacing: "-0.01em",
          }}
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
                transition: "all 0.2s ease",
                "&:hover": {
                  color: "primary.main",
                  borderColor: "primary.main",
                  transform: "translateY(-2px)",
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
                transition: "all 0.2s ease",
                "&:hover": {
                  color: "primary.main",
                  borderColor: "primary.main",
                  transform: "translateY(-2px)",
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
