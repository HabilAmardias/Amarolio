import { Container, Typography, Box, Button, Chip, Avatar } from "@mui/material";
import { useProfileController } from "../../controllers/useProfileController";
import { useProjectController } from "../../controllers/useProjectController";
import GlassBox from "../../components/common/GlassBox";
import ProjectCard from "../../components/cards/ProjectCard";
import SectionHeader from "../../components/common/SectionHeader";
import ProfilePicture from "../../assets/profile-picture.jpg";
import { Helmet } from "react-helmet-async";

export default function HomeView() {
  const { profile } = useProfileController();
  const { projects } = useProjectController();
  const featuredProjects = projects.filter((p) => p.featured).slice(0, 3);
  const domain = window.location.origin;

  return (
    <Container maxWidth="lg">
      <Helmet>
        <title>{`${profile.name} | ${profile.title}`}</title>
        <meta name="description" content={profile.bio?.substring(0, 160)} />
        <link rel="canonical" href={domain} />
        <script type="application/ld+json">
          {JSON.stringify({
            "@context": "https://schema.org",
            "@type": "Person",
            "name": profile.name,
            "url": domain,
            "description": profile.bio,
          })}
        </script>
      </Helmet>
      {/* Profile Section */}
      <GlassBox sx={{ mb: 6, textAlign: "center", py: { xs: 4, sm: 6, md: 8 }, px: { xs: 2, sm: 3, md: 4 } }}>
        <Box sx={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
          {/* Profile Picture */}
          <Avatar
            sx={{
              width: { xs: 120, sm: 150, md: 180 },
              height: { xs: 120, sm: 150, md: 180 },
              mb: 3,
              border: "3px solid",
              borderColor: "primary.main",
              boxShadow: "0 10px 30px rgba(0, 0, 0, 0.2)",
            }}
            alt={profile.name}
            src={ProfilePicture}
          >
            {profile.name?.charAt(0).toUpperCase()}
          </Avatar>

          {/* Name */}
          <Typography
            variant="h2"
            component="h1"
            sx={{
              fontSize: { xs: "1.75rem", sm: "2.25rem", md: "2.75rem" },
              mb: 2,
              color: "text.primary",
              fontWeight: 600,
            }}
          >
            {profile.name}
          </Typography>

          {/* Title */}
          <Typography
            variant="h5"
            sx={{
              fontSize: { xs: "1rem", sm: "1.25rem", md: "1.5rem" },
              mb: 3,
              color: "primary.main",
              fontWeight: 500,
            }}
          >
            {profile.title}
          </Typography>

          {/* Bio */}
          <Typography
            variant="body1"
            sx={{
              fontSize: { xs: "0.95rem", sm: "1rem", md: "1.1rem" },
              lineHeight: { xs: 1.6, md: 1.8 },
              maxWidth: 700,
              mx: "auto",
              color: "text.secondary",
              whiteSpace: "pre-wrap",
              mb: 4,
            }}
          >
            {profile.bio}
          </Typography>

          {/* CTA Button */}
          <Button variant="contained" href="/projects" size="large" sx={{ textTransform: "none", fontSize: "1rem" }}>
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
          {["TypeScript", "Go", "Python", "Docker", "Redis", "PostgreSQL"].map((skill) => (
            <Chip key={skill} label={skill} variant="outlined" />
          ))}
        </Box>
      </GlassBox>
    </Container>
  );
}