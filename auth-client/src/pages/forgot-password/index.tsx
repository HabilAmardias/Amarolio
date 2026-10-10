import { useEffect, useMemo, useState } from "react";
import {
  Alert,
  Box,
  Button,
  Paper,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { Link as RouterLink, useSearchParams } from "react-router-dom";
import MarkEmailReadOutlinedIcon from "@mui/icons-material/MarkEmailReadOutlined";
import { AuthShell } from "../../components/auth/AuthShell";
import { Seo } from "../../components/common/Seo";
import { sendResetPasswordEmail } from "../../api/auth.api";
import { toApiError } from "../../lib/errors";
import { persistRedirectUri, resolveRedirectUri } from "../../lib/redirect";
import { validateEmail } from "../../lib/validation";

export function ForgotPasswordPage() {
  const [searchParams] = useSearchParams();
  const redirectParam = searchParams.get("redirect_uri");

  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isSent, setIsSent] = useState(false);

  useEffect(() => {
    persistRedirectUri(resolveRedirectUri(redirectParam));
  }, [redirectParam]);

  const signInLink = useMemo(
    () =>
      redirectParam
        ? `/login?redirect_uri=${encodeURIComponent(redirectParam)}`
        : "/login",
    [redirectParam]
  );

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    const emailError = validateEmail(email);
    if (emailError) {
      setError(emailError);
      return;
    }

    setError("");
    setIsLoading(true);
    try {
      await sendResetPasswordEmail(email.trim());
      setIsSent(true);
    } catch (err) {
      setError(toApiError(err).message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <AuthShell subtitle="Reset password">
      <Seo title="Reset password | Amarolio Account" />
      <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
        {isSent ? (
          <Stack
            spacing={2.5}
            sx={{ alignItems: "center", textAlign: "center" }}
          >
            <Box
              sx={{
                width: 56,
                height: 56,
                borderRadius: "50%",
                display: "grid",
                placeItems: "center",
                bgcolor: "primary.main",
                color: "#fff",
              }}
            >
              <MarkEmailReadOutlinedIcon />
            </Box>
            <Box>
              <Typography variant="h6" sx={{ fontWeight: 700 }}>
                Check your inbox
              </Typography>
              <Typography variant="body2" color="text.secondary">
                If an account exists for {email.trim()}, we've sent a link to
                reset your password.
              </Typography>
            </Box>
            <Button
              fullWidth
              variant="contained"
              component={RouterLink}
              to={signInLink}
            >
              Back to sign in
            </Button>
          </Stack>
        ) : (
          <Box component="form" onSubmit={handleSubmit}>
            <Typography variant="h5" sx={{ fontWeight: 700, mb: 0.5 }}>
              Forgot your password?
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 3 }}>
              Enter your email and we'll send you a link to reset it.
            </Typography>

            {error && (
              <Alert severity="error" sx={{ mb: 2 }}>
                {error}
              </Alert>
            )}

            <TextField
              label="Email"
              type="email"
              fullWidth
              value={email}
              onChange={(event) => setEmail(event.target.value)}
              autoComplete="email"
            />

            <Button
              type="submit"
              fullWidth
              variant="contained"
              disabled={isLoading}
              sx={{ mt: 3, py: 1.4 }}
            >
              {isLoading ? "Sending..." : "Send reset link"}
            </Button>

            <Typography
              variant="body2"
              color="text.secondary"
              sx={{ mt: 3, textAlign: "center" }}
            >
              Remembered it?{" "}
              <Box
                component={RouterLink}
                to={signInLink}
                sx={{ color: "primary.main", fontWeight: 600 }}
              >
                Back to sign in
              </Box>
            </Typography>
          </Box>
        )}
      </Paper>
    </AuthShell>
  );
}
