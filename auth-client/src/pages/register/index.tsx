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
import { register, resendVerification } from "../../api/auth.api";
import { toApiError } from "../../lib/errors";
import {
  persistRedirectUri,
  resolveRedirectUri,
} from "../../lib/redirect";
import { validateEmail, validatePassword } from "../../lib/validation";

export function RegisterPage() {
  const [searchParams] = useSearchParams();
  const redirectParam = searchParams.get("redirect_uri");

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirm, setConfirm] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isDone, setIsDone] = useState(false);
  const [resendState, setResendState] = useState<"idle" | "sending" | "sent">(
    "idle"
  );

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
    const passwordError = validatePassword(password);
    if (emailError || passwordError) {
      setError(emailError ?? passwordError ?? "");
      return;
    }
    if (password !== confirm) {
      setError("Passwords do not match.");
      return;
    }

    setError("");
    setIsLoading(true);
    try {
      await register(email.trim(), password);
      setIsDone(true);
    } catch (err) {
      setError(toApiError(err).message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleResend = async () => {
    setResendState("sending");
    try {
      await resendVerification(email.trim());
      setResendState("sent");
    } catch (err) {
      setError(toApiError(err).message);
      setResendState("idle");
    }
  };

  return (
    <AuthShell subtitle="Create account">
      <Seo title="Create account | Amarolio Account" />
      <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
        {isDone ? (
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
                We sent a verification link to {email.trim()}. Confirm your
                email to activate your account.
              </Typography>
            </Box>
            <Button
              fullWidth
              variant="outlined"
              startIcon={<MarkEmailReadOutlinedIcon />}
              disabled={resendState === "sending" || resendState === "sent"}
              onClick={handleResend}
            >
              {resendState === "sent"
                ? "Verification email sent"
                : resendState === "sending"
                  ? "Sending..."
                  : "Resend verification email"}
            </Button>
            <Button fullWidth variant="contained" component={RouterLink} to={signInLink}>
              Go to sign in
            </Button>
          </Stack>
        ) : (
          <Box component="form" onSubmit={handleSubmit}>
            <Typography variant="h5" sx={{ fontWeight: 700, mb: 0.5 }}>
              Create your account
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 3 }}>
              One account for every Amarolio service.
            </Typography>

            {error && (
              <Alert severity="error" sx={{ mb: 2 }}>
                {error}
              </Alert>
            )}

            <Stack spacing={2}>
              <TextField
                label="Email"
                type="email"
                fullWidth
                value={email}
                onChange={(event) => setEmail(event.target.value)}
                autoComplete="email"
              />
              <TextField
                label="Password"
                type="password"
                fullWidth
                value={password}
                onChange={(event) => setPassword(event.target.value)}
                autoComplete="new-password"
                helperText="8–13 characters, letters and numbers only."
              />
              <TextField
                label="Confirm password"
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
              {isLoading ? "Creating account..." : "Create account"}
            </Button>

            <Typography
              variant="body2"
              color="text.secondary"
              sx={{ mt: 3, textAlign: "center" }}
            >
              Already have an account?{" "}
              <Box
                component={RouterLink}
                to={signInLink}
                sx={{ color: "primary.main", fontWeight: 600 }}
              >
                Sign in
              </Box>
            </Typography>
          </Box>
        )}
      </Paper>
    </AuthShell>
  );
}
