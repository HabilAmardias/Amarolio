import { Container, Box } from "@mui/material";
import { useExperienceController } from "../../controllers/useExperienceController";
import ExperienceCard from "../../components/cards/ExperienceCard";
import SectionHeader from "../../components/common/SectionHeader";
import GlassBox from "../../components/common/GlassBox";

export default function ExperienceView() {
  const { experiences } = useExperienceController();

  return (
    <Container maxWidth="lg">
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