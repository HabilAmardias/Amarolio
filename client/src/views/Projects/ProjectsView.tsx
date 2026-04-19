import { Container, Box, Chip } from "@mui/material";
import { useProjectController } from "../../controllers/useProjectController";
import ProjectCard from "../../components/cards/ProjectCard";
import SectionHeader from "../../components/common/SectionHeader";
import GlassBox from "../../components/common/GlassBox";

export default function ProjectsView() {
  const { projects, allTags, activeFilter, setActiveFilter } = useProjectController();

  return (
    <Container maxWidth="lg">
      <SectionHeader
        title="Projects"
        subtitle="A collection of my work and personal projects"
      />
      <GlassBox sx={{ mb: 4 }}>
        <Box sx={{ display: "flex", flexWrap: "wrap", gap: 1, justifyContent: "center" }}>
          {allTags.map((tag) => (
            <Chip
              key={tag}
              label={tag}
              onClick={() => setActiveFilter(tag)}
              color={activeFilter === tag ? "primary" : "default"}
              variant={activeFilter === tag ? "filled" : "outlined"}
              clickable
              aria-pressed={activeFilter === tag}
            />
          ))}
        </Box>
      </GlassBox>
      <Box sx={{ display: "grid", gridTemplateColumns: "repeat(auto-fit, minmax(300px, 1fr))", gap: 3 }}>
        {projects.map((project) => (
          <ProjectCard key={project.id} project={project} />
        ))}
      </Box>
    </Container>
  );
}