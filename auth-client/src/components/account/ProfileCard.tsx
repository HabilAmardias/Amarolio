import { Avatar, Box, Chip, Stack, Typography } from "@mui/material";
import VerifiedUserOutlinedIcon from "@mui/icons-material/VerifiedUserOutlined";

interface ProfileCardProps {
  username: string;
}

export function ProfileCard({ username }: ProfileCardProps) {
  return (
    <Stack
      direction={{ xs: "column", sm: "row" }}
      spacing={2.5}
      sx={{
        alignItems: { xs: "center", sm: "flex-start" },
        textAlign: { xs: "center", sm: "left" },
      }}
    >
      <Avatar
        sx={{
          width: 72,
          height: 72,
          fontSize: "1.9rem",
          fontWeight: 700,
          fontFamily: "'Poppins', 'Open Sans', sans-serif",
          background: "linear-gradient(160deg, #8FB3F0 0%, #3D6BD4 100%)",
          color: "#fff",
        }}
      >
        {username.charAt(0).toUpperCase()}
      </Avatar>

      <Box sx={{ minWidth: 0, pt: { sm: 0.5 } }}>
        <Typography
          variant="h5"
          sx={{
            fontFamily: "'Poppins', 'Open Sans', sans-serif",
            fontWeight: 700,
            wordBreak: "break-word",
          }}
        >
          {username}
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>
          @{username}
        </Typography>
        <Chip
          size="small"
          icon={<VerifiedUserOutlinedIcon />}
          label="Signed in"
          color="primary"
          variant="outlined"
        />
      </Box>
    </Stack>
  );
}
