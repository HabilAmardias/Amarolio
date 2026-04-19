import { Typography, Box } from "@mui/material";

interface SectionHeaderProps {
  title: string;
  subtitle?: string;
}

export default function SectionHeader({ title, subtitle }: SectionHeaderProps) {
  return (
    <Box sx={{ mb: 6, textAlign: "center" }}>
      <Typography variant="h2" component="h1" gutterBottom>
        {title}
      </Typography>
      {subtitle && (
        <Typography variant="body1" color="text.secondary" sx={{ maxWidth: 600, mx: "auto" }}>
          {subtitle}
        </Typography>
      )}
    </Box>
  );
}