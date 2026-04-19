import { Container, Typography, Box, Button, Chip } from "@mui/material";
import { useProfileController } from "../../controllers/useProfileController";
import { useProjectController } from "../../controllers/useProjectController";
import GlassBox from "../../components/common/GlassBox";
import ProjectCard from "../../components/cards/ProjectCard";
import SectionHeader from "../../components/common/SectionHeader";

export default function HomeView() {
  const { profile } = useProfileController();
  const { projects } = useProjectController();
  const featuredProjects = projects.filter((p) => p.featured).slice(0, 3);

  return (
    <Container maxWidth="lg">
      <GlassBox sx={{ mb: 6, textAlign: "center" }}>
        <Typography variant="h1" component="h1" gutterBottom>
          {profile.name} {/* TODO: Replace with real content */}
        </Typography>
        <Typography variant="h4" color="primary" gutterBottom>
          {profile.title} {/* TODO: Replace with real content */}
        </Typography>
        <Typography variant="body1" color="text.secondary" sx={{ mb: 4, maxWidth: 600, mx: "auto" }}>
          {profile.bio} {/* TODO: Replace with real content */}
        </Typography>
        <Box sx={{ display: "flex", gap: 2, justifyContent: "center" }}>
          <Button variant="contained" href="/projects" size="large">
            View Projects
          </Button>
        </Box>
      </GlassBox>

      <SectionHeader title="Featured Projects" subtitle="A selection of my recent work" />
      <Box sx={{ display: "grid", gridTemplateColumns: "repeat(auto-fit, minmax(300px, 1fr))", gap: 3, mb: 6 }}>
        {featuredProjects.map((project) => (
          <ProjectCard key={project.id} project={project} />
        ))}
      </Box>

      <GlassBox>
        <Typography variant="h3" gutterBottom>
          Skills & Technologies
        </Typography>
        <Box sx={{ display: "flex", flexWrap: "wrap", gap: 1 }}>
          {[ "TypeScript", "Go", "Python", "Docker", "Redis", "PostgreSQL"].map((skill) => (
            <Chip key={skill} label={skill} variant="outlined" />
          ))}
        </Box>
      </GlassBox>
    </Container>
  );
}