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
import CheckCircleOutlinedIcon from "@mui/icons-material/CheckCircleOutlined";
import { AuthShell } from "../../components/auth/AuthShell";
import { Seo } from "../../components/common/Seo";
import { resetPassword, sendResetPasswordEmail } from "../../api/auth.api";
import { toApiError } from "../../lib/errors";
import {
  clearPersistedRedirectUri,
  persistRedirectUri,
  resolveRedirectUri,
} from "../../lib/redirect";
import { validatePassword } from "../../lib/validation";

export function ResetPasswordPage() {
  const [searchParams] = useSearchParams();
  const userId = searchParams.get("user_id") ?? "";
  const token = searchParams.get("token") ?? "";
  const redirectParam = searchParams.get("redirect_uri");

  const hasParams = Boolean(userId && token);

  const [newPassword, setNewPassword] = useState("");
  const [confirm, setConfirm] = useState("");
  const [error, setError] = useState(
    hasParams ? "" : "This reset link is incomplete or invalid."
  );
  const [isLoading, setIsLoading] = useState(false);
  const [isDone, setIsDone] = useState(false);
  const [showRequestLink, setShowRequestLink] = useState(!hasParams);
  const [email, setEmail] = useState("");
  const [resendState, setResendState] = useState<"idle" | "sending" | "sent">(
    "idle"
  );
  const [resendError, setResendError] = useState("");

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

    const passwordError = validatePassword(newPassword);
    if (passwordError) {
      setError(passwordError);
      return;
    }
    if (newPassword !== confirm) {
      setError("Passwords do not match.");
      return;
    }

    setError("");
    setIsLoading(true);
    try {
      await resetPassword(userId, token, newPassword);
      clearPersistedRedirectUri();
      setIsDone(true);
    } catch (err) {
      const apiError = toApiError(err);
      setError(apiError.message);
      if (apiError.code === 40101 || apiError.status === 401) {
        setShowRequestLink(true);
      }
    } finally {
      setIsLoading(false);
    }
  };

  const handleRequestLink = async () => {
    setResendError("");
    setResendState("sending");
    try {
      await sendResetPasswordEmail(email.trim());
      setResendState("sent");
    } catch (err) {
      setResendError(toApiError(err).message);
      setResendState("idle");
    }
  };

  return (
    <AuthShell subtitle="Reset password">
      <Seo title="Reset password | Amarolio Account" />
      <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
        {isDone ? (
          <Stack
            spacing={2.5}
            sx={{ alignItems: "center", textAlign: "center" }}
          >
            <Box sx={{ color: "success.main" }}>
              <CheckCircleOutlinedIcon sx={{ fontSize: 56 }} />
            </Box>
            <Box>
              <Typography variant="h6" sx={{ fontWeight: 700 }}>
                Password updated
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Your password has been changed. Sign in with your new password.
              </Typography>
            </Box>
            <Button
              fullWidth
              variant="contained"
              component={RouterLink}
              to={signInLink}
              sx={{ py: 1.4 }}
            >
              Go to sign in
            </Button>
          </Stack>
        ) : (
          <Stack spacing={2.5}>
            <Box>
              <Typography variant="h5" sx={{ fontWeight: 700, mb: 0.5 }}>
                Choose a new password
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Pick a strong password for your Amarolio account.
              </Typography>
            </Box>

            {error && <Alert severity="error">{error}</Alert>}

            {hasParams && (
              <Box component="form" onSubmit={handleSubmit}>
                <Stack spacing={2}>
                  <TextField
                    label="New password"
                    type="password"
                    fullWidth
                    value={newPassword}
                    onChange={(event) => setNewPassword(event.target.value)}
                    autoComplete="new-password"
                    helperText="8–13 characters, letters and numbers only."
                  />
                  <TextField
                    label="Confirm new password"
                    type="password"
                    fullWidth
                    value={confirm}
                    onChange={(event) => setConfirm(event.target.value)}
                    autoComplete="new-password"
                  />
                </Stack>

                <Button
                  type="submit"
                  fullWidth
                  variant="contained"
                  disabled={isLoading}
                  sx={{ mt: 3, py: 1.4 }}
                >
                  {isLoading ? "Resetting..." : "Reset password"}
                </Button>
              </Box>
            )}

            {showRequestLink && (
              <Box>
                <Typography
                  variant="body2"
                  color="text.secondary"
                  sx={{ mb: 1 }}
                >
                  Request a new reset link:
                </Typography>
                {resendState === "sent" ? (
                  <Alert severity="success">
                    If an account exists for that address, we've sent a reset
                    link.
                  </Alert>
                ) : (
                  <Stack
                    direction={{ xs: "column", sm: "row" }}
                    spacing={1}
                    sx={{ alignItems: { sm: "flex-start" } }}
                  >
                    <TextField
                      label="Email"
                      type="email"
                      size="small"
                      fullWidth
                      value={email}
                      onChange={(event) => setEmail(event.target.value)}
                    />
                    <Button
                      variant="outlined"
                      disabled={resendState === "sending" || !email}
                      onClick={handleRequestLink}
                      sx={{ flexShrink: 0, height: 40 }}
                    >
                      {resendState === "sending" ? "Sending..." : "Send"}
                    </Button>
                  </Stack>
                )}
                {resendError && (
                  <Alert severity="error" sx={{ mt: 1 }}>
                    {resendError}
                  </Alert>
                )}
              </Box>
            )}

            <Button variant="text" component={RouterLink} to={signInLink}>
              Back to sign in
            </Button>
          </Stack>
        )}
      </Paper>
    </AuthShell>
  );
}
