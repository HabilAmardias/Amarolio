import { Container, Typography, Box, Button, Chip, Avatar } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import { NavLink } from "react-router-dom";
import ArrowForward from "@mui/icons-material/ArrowForward";
import { useProfileController } from "../../controllers/useProfileController";
import { useProjectController } from "../../controllers/useProjectController";
import { useExperienceController } from "../../controllers/useExperienceController";
import GlassBox from "../../components/common/GlassBox";
import ProjectCard from "../../components/cards/ProjectCard";
import SectionHeader from "../../components/common/SectionHeader";
import Seo from "../../components/common/Seo";
import ProfilePicture from "../../assets/profile-picture.jpg";
import { siteConfig } from "../../config/site";

const skills = ["TypeScript", "Go", "Python", "Docker", "Redis", "PostgreSQL"];

export default function HomeView() {
  const theme = useTheme();
  const isLight = theme.palette.mode === "light";
  const { profile } = useProfileController();
  const { projects } = useProjectController();
  const { experiences } = useExperienceController();
  const featuredProjects = projects.filter((p) => p.featured).slice(0, 3);

  const stats = [
    { value: experiences.length, label: "Roles held" },
    { value: projects.length, label: "Projects built" },
    { value: skills.length, label: "Technologies" },
  ];

  return (
    <Container maxWidth="lg" sx={{ py: { xs: 3, md: 5 } }}>
      <Seo
        title={profile.name || siteConfig.name}
        description={profile.bio?.substring(0, 160) || siteConfig.description}
        path="/"
        type="profile"
        image={ProfilePicture}
        keywords={[
          profile.name,
          "software engineer",
          "portfolio",
          "web development",
          "Amarolio",
        ].filter(Boolean)}
        jsonLd={[
          {
            "@context": "https://schema.org",
            "@type": "Person",
            "name": profile.name,
            "url": siteConfig.url,
            "description": profile.bio,
            "jobTitle": profile.title,
          },
          {
            "@context": "https://schema.org",
            "@type": "WebSite",
            "name": siteConfig.name,
            "url": siteConfig.url,
            "inLanguage": "en",
          },
        ]}
      />

      <GlassBox
        sx={{
          mb: 8,
          py: { xs: 5, md: 7 },
          px: { xs: 3, sm: 4, md: 6 },
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: { xs: "column", md: "row" },
            alignItems: "center",
            gap: { xs: 4, md: 6 },
          }}
        >
          {/* Avatar */}
          <Box sx={{ order: { xs: 0, md: 2 }, flexShrink: 0 }}>
            <Box
              sx={{
                width: { xs: 148, sm: 176, md: 200 },
                height: { xs: 148, sm: 176, md: 200 },
                borderRadius: "50%",
                p: "5px",
                background: isLight
                  ? "linear-gradient(135deg, #384959, #88bdf2)"
                  : "linear-gradient(135deg, #bdddfc, #6a89a7)",
                boxShadow: isLight
                  ? "0 16px 40px -12px rgba(56,73,89,0.45)"
                  : "0 16px 40px -12px rgba(0,0,0,0.5)",
              }}
            >
              <Avatar
                sx={{
                  width: "100%",
                  height: "100%",
                  border: "4px solid",
                  borderColor: "background.paper",
                  fontSize: "3rem",
                  fontWeight: 700,
                }}
                alt={profile.name}
                src={ProfilePicture}
              >
                {profile.name?.charAt(0).toUpperCase()}
              </Avatar>
            </Box>
          </Box>

          {/* Text */}
          <Box
            sx={{
              flex: 1,
              minWidth: 0,
              order: { xs: 1, md: 0 },
              textAlign: { xs: "center", md: "left" },
            }}
          >

            <Typography
              variant="h1"
              component="h1"
              sx={{
                fontSize: { xs: "2rem", sm: "2.6rem", md: "3rem" },
                mb: 1.25,
              }}
            >
              {profile.name}
            </Typography>

            <Typography
              variant="h5"
              sx={{
                fontSize: { xs: "1.1rem", sm: "1.3rem" },
                mb: 2.5,
                fontWeight: 700,
                background: isLight
                  ? "linear-gradient(90deg, #384959, #6a89a7, #88bdf2)"
                  : "linear-gradient(90deg, #bdddfc, #88bdf2)",
                WebkitBackgroundClip: "text",
                WebkitTextFillColor: "transparent",
                backgroundClip: "text",
              }}
            >
              {profile.title}
            </Typography>

            <Typography
              variant="body1"
              sx={{
                fontSize: { xs: "1rem", md: "1.1rem" },
                lineHeight: 1.8,
                maxWidth: 620,
                mx: { xs: "auto", md: 0 },
                color: "text.secondary",
                whiteSpace: "pre-wrap",
                mb: 3.5,
              }}
            >
              {profile.bio}
            </Typography>

            <Box
              sx={{
                display: "flex",
                gap: 1.5,
                flexWrap: "wrap",
                justifyContent: { xs: "center", md: "flex-start" },
              }}
            >
              <Button
                variant="contained"
                component={NavLink}
                to="/projects"
                size="large"
                endIcon={<ArrowForward fontSize="small" />}
                sx={{ px: 3.5, py: 1.25 }}
              >
                View Projects
              </Button>
              <Button
                variant="outlined"
                component={NavLink}
                to="/experience"
                size="large"
                sx={{ px: 3.5, py: 1.25 }}
              >
                My Experience
              </Button>
            </Box>

            <Box
              sx={{
                display: "flex",
                flexWrap: "wrap",
                gap: { xs: 2.5, sm: 4 },
                justifyContent: { xs: "center", md: "flex-start" },
                mt: 3.5,
              }}
            >
              {stats.map((stat) => (
                <Box key={stat.label}>
                  <Typography
                    variant="h5"
                    component="div"
                    sx={{ fontWeight: 800, lineHeight: 1.1 }}
                  >
                    {stat.value}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {stat.label}
                  </Typography>
                </Box>
              ))}
            </Box>
          </Box>
        </Box>
      </GlassBox>

      <SectionHeader
        eyebrow="My Work"
        title="Featured Projects"
        subtitle="A selection of my recent projects"
      />
      <Box
        sx={{
          display: "grid",
          gridTemplateColumns:
            "repeat(auto-fit, minmax(min(300px, 100%), 1fr))",
          gap: 3,
          mb: 8,
        }}
      >
        {featuredProjects.map((project) => (
          <ProjectCard key={project.id} project={project} />
        ))}
      </Box>

      <GlassBox>
        <SectionHeader
          eyebrow="Toolkit"
          title="Skills & Technologies"
          subtitle="Technologies I work with across my projects"
        />
        <Box
          sx={{
            display: "flex",
            flexWrap: "wrap",
            gap: 1,
            justifyContent: { xs: "center", sm: "flex-start" },
          }}
        >
          {skills.map((skill) => (
            <Chip
              key={skill}
              label={skill}
              variant="outlined"
              sx={{ px: 0.5, py: 0.5 }}
            />
          ))}
        </Box>
      </GlassBox>
    </Container>
  );
}
