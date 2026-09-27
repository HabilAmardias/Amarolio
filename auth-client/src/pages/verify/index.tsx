import { useEffect, useMemo, useRef, useState } from "react";
import {
  Alert,
  Box,
  Button,
  CircularProgress,
  Paper,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { Link as RouterLink, useSearchParams } from "react-router-dom";
import CheckCircleOutlinedIcon from "@mui/icons-material/CheckCircleOutlined";
import ErrorOutlinedIcon from "@mui/icons-material/ErrorOutlined";
import { AuthShell } from "../../components/auth/AuthShell";
import { Seo } from "../../components/common/Seo";
import { resendVerification, verify } from "../../api/auth.api";
import { toApiError } from "../../lib/errors";
import {
  persistRedirectUri,
  resolveRedirectUri,
} from "../../lib/redirect";

type Status = "verifying" | "success" | "error";

export function VerifyPage() {
  const [searchParams] = useSearchParams();
  const userId = searchParams.get("user_id") ?? "";
  const token = searchParams.get("token") ?? "";
  const redirectParam = searchParams.get("redirect_uri");

  const hasParams = Boolean(userId && token);
  const [status, setStatus] = useState<Status>(hasParams ? "verifying" : "error");
  const [error, setError] = useState(
    hasParams ? "" : "This verification link is incomplete or invalid."
  );
  const [email, setEmail] = useState("");
  const [resendState, setResendState] = useState<"idle" | "sending" | "sent">(
    "idle"
  );
  const [resendError, setResendError] = useState("");
  const ran = useRef(false);

  useEffect(() => {
    persistRedirectUri(resolveRedirectUri(redirectParam));
  }, [redirectParam]);

  useEffect(() => {
    if (ran.current || !hasParams) return;
    ran.current = true;

    verify(userId, token)
      .then(() => setStatus("success"))
      .catch((err) => {
        setStatus("error");
        setError(toApiError(err).message);
      });
  }, [userId, token, hasParams]);

  const signInLink = useMemo(
    () =>
      redirectParam
        ? `/login?redirect_uri=${encodeURIComponent(redirectParam)}`
        : "/login",
    [redirectParam]
  );

  const handleResend = async () => {
    setResendError("");
    setResendState("sending");
    try {
      await resendVerification(email.trim());
      setResendState("sent");
    } catch (err) {
      setResendError(toApiError(err).message);
      setResendState("idle");
    }
  };

  return (
    <AuthShell subtitle="Verify email">
      <Seo title="Verify email | Amarolio Account" />
      <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
        {status === "verifying" && (
          <Stack
            spacing={2.5}
            sx={{ alignItems: "center", textAlign: "center" }}
          >
            <CircularProgress />
            <Typography variant="body1" color="text.secondary">
              Verifying your email...
            </Typography>
          </Stack>
        )}

        {status === "success" && (
          <Stack
            spacing={2.5}
            sx={{ alignItems: "center", textAlign: "center" }}
          >
            <Box sx={{ color: "success.main" }}>
              <CheckCircleOutlinedIcon sx={{ fontSize: 56 }} />
            </Box>
            <Box>
              <Typography variant="h6" sx={{ fontWeight: 700 }}>
                Email verified
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Your Amarolio account is active. You can sign in now.
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
        )}

        {status === "error" && (
          <Stack spacing={2.5}>
            <Stack
              spacing={1.5}
              sx={{ alignItems: "center", textAlign: "center" }}
            >
              <Box sx={{ color: "error.main" }}>
                <ErrorOutlinedIcon sx={{ fontSize: 56 }} />
              </Box>
              <Box>
                <Typography variant="h6" sx={{ fontWeight: 700 }}>
                  We couldn't verify your email
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {error}
                </Typography>
              </Box>
            </Stack>

            <Box>
              <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>
                Request a new verification link:
              </Typography>
              {resendState === "sent" ? (
                <Alert severity="success">
                  Verification email sent. Check your inbox.
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
                    onClick={handleResend}
                    sx={{ flexShrink: 0, height: 40 }}
                  >
                    {resendState === "sending" ? "Sending..." : "Resend"}
                  </Button>
                </Stack>
              )}
              {resendError && (
                <Alert severity="error" sx={{ mt: 1 }}>
                  {resendError}
                </Alert>
              )}
            </Box>

            <Button variant="text" component={RouterLink} to={signInLink}>
              Go to sign in
            </Button>
          </Stack>
        )}
      </Paper>
    </AuthShell>
  );
}
