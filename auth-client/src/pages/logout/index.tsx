import { useEffect } from "react";
import { Box, CircularProgress, Typography } from "@mui/material";
import { useSearchParams } from "react-router-dom";
import { AuthShell } from "../../components/auth/AuthShell";
import { Seo } from "../../components/common/Seo";
import { logout } from "../../api/auth.api";
import {
  clearPersistedRedirectUri,
  resolveRedirectUri,
} from "../../lib/redirect";

export function LogoutPage() {
  const [searchParams] = useSearchParams();
  const redirectParam = searchParams.get("redirect_uri");

  useEffect(() => {
    let target = resolveRedirectUri(redirectParam);
    if (target === "/") target = "/login";

    logout().finally(() => {
      clearPersistedRedirectUri();
      window.location.href = target;
    });
  }, [redirectParam]);

  return (
    <AuthShell subtitle="Sign out">
      <Seo title="Signing out | Amarolio Account" />
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: 2,
          py: 8,
        }}
      >
        <CircularProgress />
        <Typography variant="body1" color="text.secondary">
          Signing you out...
        </Typography>
      </Box>
    </AuthShell>
  );
}
