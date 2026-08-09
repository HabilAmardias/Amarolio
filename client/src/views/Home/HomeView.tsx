import { Container, Typography, Box, Button, Chip, Avatar } from "@mui/material";
import { useTheme } from "@mui/material/styles";
import { NavLink } from "react-router-dom";
import ArrowForward from "@mui/icons-material/ArrowForward";
import Terminal from "@mui/icons-material/Terminal";
import LocationOn from "@mui/icons-material/LocationOn";
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

      {/* Hero */}
      <Box sx={{ position: "relative", mb: { xs: 7, md: 9 } }}>
        <Box
          className="blob"
          sx={{
            width: { xs: 260, md: 380 },
            height: { xs: 260, md: 380 },
            top: -80,
            right: { xs: -120, md: -60 },
            background: isLight
              ? "rgba(143,179,240,0.35)"
              : "rgba(143,179,240,0.14)",
          }}
        />
        <Box
          className="blob"
          sx={{
            width: { xs: 200, md: 300 },
            height: { xs: 200, md: 300 },
            bottom: -60,
            left: -140,
            background: isLight
              ? "rgba(240,166,59,0.22)"
              : "rgba(240,166,59,0.1)",
            animationDelay: "-6s",
          }}
        />

        <Box
          sx={{
            position: "relative",
            display: "flex",
            flexDirection: { xs: "column", md: "row" },
            alignItems: "center",
            gap: { xs: 5, md: 7 },
            py: { xs: 4, md: 5 },
          }}
        >
          {/* Text */}
          <Box
            className="animate-fade-up"
            sx={{
              flex: 1,
              minWidth: 0,
              textAlign: { xs: "center", md: "left" },
            }}
          >

            <Typography
              variant="h1"
              component="h1"
              sx={{
                fontSize: { xs: "2.1rem", sm: "2.8rem", md: "3.2rem" },
                mb: 1.25,
              }}
            >
              Hey, I&apos;m{" "}
              <Box
                component="span"
                sx={{
                  position: "relative",
                  display: "inline-block",
                  color: "primary.main",
                  "&::after": {
                    content: '""',
                    position: "absolute",
                    left: 0,
                    right: 0,
                    bottom: "0.06em",
                    height: "0.32em",
                    borderRadius: 999,
                    background: isLight
                      ? "rgba(240,166,59,0.35)"
                      : "rgba(240,166,59,0.3)",
                    zIndex: -1,
                  },
                }}
              >
                {profile.name}
              </Box>
            </Typography>

            <Typography
              variant="h5"
              sx={{
                fontSize: { xs: "1.15rem", sm: "1.35rem" },
                mb: 2.5,
                fontWeight: 600,
                fontFamily: "'Poppins', sans-serif",
                background: isLight
                  ? "linear-gradient(90deg, #3D6BD4, #8FB3F0, #F0A63B)"
                  : "linear-gradient(90deg, #C6DCFA, #8FB3F0, #F5B45C)",
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
                gap: { xs: 2, sm: 3 },
                justifyContent: { xs: "center", md: "flex-start" },
                mt: 4,
              }}
            >
              {stats.map((stat) => (
                <Box
                  key={stat.label}
                  sx={{
                    px: { xs: 2, sm: 2.5 },
                    py: 1.5,
                    borderRadius: "18px",
                    border: "1px solid",
                    borderColor: "divider",
                    background: isLight
                      ? "rgba(255,255,255,0.65)"
                      : "rgba(143,179,240,0.07)",
                    transition:
                      "transform 0.2s cubic-bezier(0.22,1,0.36,1), box-shadow 0.2s ease",
                    "&:hover": {
                      transform: "translateY(-3px)",
                      boxShadow: isLight
                        ? "0 12px 24px -12px rgba(37,51,80,0.25)"
                        : "0 12px 24px -10px rgba(0,0,0,0.5)",
                    },
                  }}
                >
                  <Typography
                    variant="h6"
                    component="div"
                    sx={{ fontWeight: 700, lineHeight: 1.1, fontFamily: "'Poppins', sans-serif" }}
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

          {/* Avatar */}
          <Box
            className="animate-fade-up"
            sx={{ flexShrink: 0, animationDelay: "0.15s", position: "relative" }}
          >
            <Box
              className="animate-float"
              sx={{
                position: "absolute",
                top: { xs: -28, md: -40 },
                right: { xs: -8, md: -28 },
                display: "flex",
                alignItems: "center",
                gap: 0.75,
                px: 1.5,
                py: 0.8,
                borderRadius: 999,
                background: isLight ? "#ffffff" : "#232B3A",
                border: "1px solid",
                borderColor: "divider",
                boxShadow: isLight
                  ? "0 12px 28px -12px rgba(37,51,80,0.3)"
                  : "0 12px 28px -10px rgba(0,0,0,0.5)",
              }}
            >
              <Terminal fontSize="small" sx={{ color: "warning.main" }} />
              <Typography variant="body2" sx={{ fontWeight: 600 }}>
                {skills[0]}
              </Typography>
            </Box>

            <Box
              sx={{
                position: "relative",
                width: { xs: 168, sm: 200, md: 232 },
                height: { xs: 168, sm: 200, md: 232 },
                borderRadius: "34% 66% 62% 38% / 44% 40% 60% 56%",
                background: isLight
                  ? "linear-gradient(135deg, #3D6BD4, #8FB3F0, #F0A63B)"
                  : "linear-gradient(135deg, #C6DCFA, #8FB3F0, #F5B45C)",
                padding: "6px",
                boxShadow: isLight
                  ? "0 24px 56px -18px rgba(61,107,212,0.5)"
                  : "0 24px 56px -16px rgba(0,0,0,0.55)",
                animation: "float 7s ease-in-out infinite",
              }}
            >
              <Box
                sx={{
                  position: "absolute",
                  inset: 10,
                  borderRadius: "inherit",
                  border: "2px dashed",
                  borderColor: "rgba(255,255,255,0.6)",
                }}
              />
              <Avatar
                sx={{
                  width: "100%",
                  height: "100%",
                  borderRadius: "inherit",
                  fontSize: "3.2rem",
                  fontWeight: 700,
                }}
                alt={profile.name}
                src={ProfilePicture}
              >
                {profile.name?.charAt(0).toUpperCase()}
              </Avatar>
            </Box>

            <Box
              className="animate-float"
              sx={{
                position: "absolute",
                bottom: { xs: -18, md: -26 },
                left: { xs: -12, md: -30 },
                display: "flex",
                alignItems: "center",
                gap: 0.75,
                px: 1.5,
                py: 0.8,
                borderRadius: 999,
                background: isLight ? "#ffffff" : "#232B3A",
                border: "1px solid",
                borderColor: "divider",
                boxShadow: isLight
                  ? "0 12px 28px -12px rgba(37,51,80,0.3)"
                  : "0 12px 28px -10px rgba(0,0,0,0.5)",
                animationDelay: "-3s",
              }}
            >
              <LocationOn fontSize="small" sx={{ color: "primary.main" }} />
              <Typography variant="body2" sx={{ fontWeight: 600 }}>
                Jakarta, ID
              </Typography>
            </Box>
          </Box>
        </Box>
      </Box>

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
              sx={{
                px: 0.5,
                py: 0.5,
                transition: "all 0.2s ease",
                "&:hover": {
                  transform: "translateY(-2px)",
                  bgcolor: isLight
                    ? "rgba(61,107,212,0.1)"
                    : "rgba(143,179,240,0.14)",
                },
              }}
            />
          ))}
        </Box>
      </GlassBox>
    </Container>
  );
}
