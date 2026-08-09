import { Container, Box, Chip } from "@mui/material";
import { useProjectController } from "../../controllers/useProjectController";
import ProjectCard from "../../components/cards/ProjectCard";
import SectionHeader from "../../components/common/SectionHeader";
import GlassBox from "../../components/common/GlassBox";
import Seo from "../../components/common/Seo";
import { siteConfig } from "../../config/site";

export default function ProjectsView() {
  const { projects, allTags, activeFilter, setActiveFilter } = useProjectController();

  return (
    <Container maxWidth="lg" sx={{ py: { xs: 3, md: 5 } }}>
      <Seo
        title="Projects"
        description="A curated collection of my software development projects, open-source contributions, and technical experiments."
        path="/projects"
        keywords={[
          "software projects",
          "portfolio",
          "github",
          "programming",
          "coding projects",
          "web development",
        ]}
        jsonLd={[
          {
            "@context": "https://schema.org",
            "@type": "BreadcrumbList",
            "itemListElement": [
              {
                "@type": "ListItem",
                "position": 1,
                "name": "Home",
                "item": `${siteConfig.url}/`,
              },
              {
                "@type": "ListItem",
                "position": 2,
                "name": "Projects",
                "item": `${siteConfig.url}/projects`,
              },
            ],
          },
          {
            "@context": "https://schema.org",
            "@type": "ItemList",
            "itemListElement": projects.map((project, index) => ({
              "@type": "ListItem",
              "position": index + 1,
              "name": project.title,
              "url":
                project.liveUrl && project.liveUrl !== "#"
                  ? project.liveUrl
                  : `${siteConfig.url}/projects`,
              "image": project.imageUrl,
            })),
          },
        ]}
      />
      <SectionHeader
        eyebrow="Portfolio"
        title="Projects"
        subtitle="A collection of my work and personal projects"
      />
      <GlassBox sx={{ mb: 4, p: { xs: 2, sm: 3 } }}>
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
      <Box
        sx={{
          display: "grid",
          gridTemplateColumns:
            "repeat(auto-fit, minmax(min(300px, 100%), 1fr))",
          gap: 3,
        }}
      >
        {projects.map((project) => (
          <ProjectCard key={project.id} project={project} />
        ))}
      </Box>
    </Container>
  );
}