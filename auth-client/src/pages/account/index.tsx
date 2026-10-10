import { useState } from "react";
import {
  Box,
  Button,
  CircularProgress,
  Divider,
  Paper,
  Stack,
  Typography,
} from "@mui/material";
import LogoutIcon from "@mui/icons-material/Logout";
import { Navigate, useNavigate } from "react-router-dom";
import { AuthShell } from "../../components/auth/AuthShell";
import { Seo } from "../../components/common/Seo";
import { ProfileCard } from "../../components/account/ProfileCard";
import { AppCard } from "../../components/account/AppCard";
import { useAuth } from "../../controllers/useAuth";
import { logout } from "../../api/auth.api";
import { clearPersistedRedirectUri } from "../../lib/redirect";
import { apps } from "../../config/apps";

export function AccountPage() {
  const { user, isLoading } = useAuth();
  const navigate = useNavigate();
  const [isLoggingOut, setIsLoggingOut] = useState(false);

  if (isLoading) {
    return (
      <AuthShell subtitle="Account">
        <Box sx={{ display: "grid", placeItems: "center", py: 8 }}>
          <CircularProgress />
        </Box>
      </AuthShell>
    );
  }

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  const handleLogout = async () => {
    setIsLoggingOut(true);
    await logout();
    clearPersistedRedirectUri();
    navigate("/login", { replace: true });
  };

  return (
    <AuthShell subtitle="Account">
      <Seo title="Your account | Amarolio Account" />
      <Stack spacing={3}>
        <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
          <ProfileCard username={user.username} />
          <Divider sx={{ my: 3 }} />
          <Button
            fullWidth
            variant="outlined"
            color="error"
            startIcon={<LogoutIcon />}
            disabled={isLoggingOut}
            onClick={handleLogout}
          >
            {isLoggingOut ? "Signing out..." : "Log out"}
          </Button>
        </Paper>

        <Box>
          <Typography
            variant="subtitle2"
            sx={{
              fontWeight: 700,
              textTransform: "uppercase",
              letterSpacing: "0.08em",
              color: "text.secondary",
              mb: 1.5,
            }}
          >
            Your apps
          </Typography>
          <Stack spacing={1.5}>
            {apps.map((app) => (
              <AppCard key={app.id} app={app} />
            ))}
          </Stack>
        </Box>
      </Stack>
    </AuthShell>
  );
}
