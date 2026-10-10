import { Box, Button, Paper, Stack, Typography } from "@mui/material";
import OpenInNewIcon from "@mui/icons-material/OpenInNew";
import type { AppLink } from "../../config/apps";

interface AppCardProps {
  app: AppLink;
}

export function AppCard({ app }: AppCardProps) {
  return (
    <Paper
      variant="outlined"
      sx={{
        p: 2.5,
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        gap: 2,
        transition: "border-color 0.28s ease, transform 0.28s ease",
        "&:hover": {
          borderColor: "primary.main",
          transform: "translateY(-2px)",
        },
      }}
    >
      <Stack
        direction="row"
        spacing={2}
        sx={{ alignItems: "center", minWidth: 0 }}
      >
        <Box
          aria-hidden
          sx={{
            width: 42,
            height: 42,
            flexShrink: 0,
            borderRadius: 2.5,
            display: "grid",
            placeItems: "center",
            fontWeight: 800,
            color: "#fff",
            fontFamily: "'Poppins', 'Open Sans', sans-serif",
            background: app.accent,
          }}
        >
          {app.name.charAt(0)}
        </Box>
        <Box sx={{ minWidth: 0 }}>
          <Typography sx={{ fontWeight: 700 }}>{app.name}</Typography>
          <Typography variant="body2" color="text.secondary">
            {app.description}
          </Typography>
        </Box>
      </Stack>

      <Button
        href={app.url}
        target="_blank"
        rel="noopener noreferrer"
        endIcon={<OpenInNewIcon />}
        variant="outlined"
        size="small"
        sx={{ flexShrink: 0 }}
      >
        Open
      </Button>
    </Paper>
  );
}
