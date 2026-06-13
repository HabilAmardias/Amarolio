import { Container, Box } from "@mui/material";
import { useExperienceController } from "../../controllers/useExperienceController";
import ExperienceCard from "../../components/cards/ExperienceCard";
import SectionHeader from "../../components/common/SectionHeader";
import GlassBox from "../../components/common/GlassBox";
import { Helmet } from "react-helmet-async";

export default function ExperienceView() {
  const { experiences } = useExperienceController();

  return (
    <Container maxWidth="lg">
      <Helmet>
        <title>Experience | Professional Journey</title>
        <meta name="description" content="Explore my professional journey, career highlights, and the companies I've worked with throughout my software engineering career." />
        <meta name="keywords" content="professional experience, work history, resume, career highlights, software engineer experience" />
      </Helmet>
      <SectionHeader
        title="Experience"
        subtitle="My professional journey and career highlights"
      />
      <GlassBox>
        <Box sx={{ display: "flex", flexDirection: "column", gap: 3 }}>
          {experiences.map((experience) => (
            <ExperienceCard key={experience.id} experience={experience} />
          ))}
        </Box>
      </GlassBox>
    </Container>
  );
}