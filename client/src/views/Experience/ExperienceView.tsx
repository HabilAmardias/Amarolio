import { Container, Box } from "@mui/material";
import { useExperienceController } from "../../controllers/useExperienceController";
import ExperienceCard from "../../components/cards/ExperienceCard";
import SectionHeader from "../../components/common/SectionHeader";
import GlassBox from "../../components/common/GlassBox";
import Seo from "../../components/common/Seo";
import { siteConfig } from "../../config/site";

export default function ExperienceView() {
  const { experiences } = useExperienceController();

  return (
    <Container maxWidth="lg" sx={{ py: { xs: 3, md: 5 } }}>
      <Seo
        title="Experience"
        description="Explore my professional journey, career highlights, and the companies I've worked with throughout my software engineering career."
        path="/experience"
        keywords={[
          "professional experience",
          "work history",
          "resume",
          "career highlights",
          "software engineer experience",
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
                "name": "Experience",
                "item": `${siteConfig.url}/experience`,
              },
            ],
          },
        ]}
      />
      <SectionHeader
        eyebrow="Career"
        title="Experience"
        subtitle="My professional journey and career highlights"
      />
      <GlassBox sx={{ p: { xs: 2.5, sm: 4 } }}>
        <Box sx={{ display: "flex", flexDirection: "column", gap: 3 }}>
          {experiences.map((experience) => (
            <ExperienceCard key={experience.id} experience={experience} />
          ))}
        </Box>
      </GlassBox>
    </Container>
  );
}
