import { useEffect, useMemo, useState } from "react";
import {
  Alert,
  Avatar,
  Box,
  Button,
  CircularProgress,
  Divider,
  Paper,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { Link as RouterLink, useSearchParams } from "react-router-dom";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import MarkEmailReadOutlinedIcon from "@mui/icons-material/MarkEmailReadOutlined";
import { AuthShell } from "../../components/auth/AuthShell";
import { OtpInput } from "../../components/auth/OtpInput";
import { ResendButton } from "../../components/auth/ResendButton";
import { Seo } from "../../components/common/Seo";
import { useLoginFlow } from "../../controllers/useLoginFlow";
import { useAuth } from "../../controllers/useAuth";
import { resendVerification } from "../../api/auth.api";
import {
  getPersistedRedirectUri,
  persistRedirectUri,
  resolveRedirectUri,
} from "../../lib/redirect";
import { toApiError } from "../../lib/errors";

export function LoginPage() {
  const [searchParams] = useSearchParams();
  const redirectParam = searchParams.get("redirect_uri");
  const emailParam = searchParams.get("email") ?? "";

  const { user, isLoading: isAuthLoading } = useAuth();
  const [forceForm, setForceForm] = useState(false);
  const flow = useLoginFlow(emailParam);

  const [resendState, setResendState] = useState<"idle" | "sending" | "sent">(
    "idle"
  );
  const [resendError, setResendError] = useState("");

  useEffect(() => {
    persistRedirectUri(resolveRedirectUri(redirectParam));
  }, [redirectParam]);

  const registerLink = useMemo(
    () =>
      redirectParam
        ? `/register?redirect_uri=${encodeURIComponent(redirectParam)}`
        : "/register",
    [redirectParam]
  );

  const showUnverified =
    flow.errorCode === 40001 && /not verified/i.test(flow.error);

  const handleResendVerification = async () => {
    setResendError("");
    setResendState("sending");
    try {
      await resendVerification(flow.email.trim());
      setResendState("sent");
    } catch (err) {
      setResendError(toApiError(err).message);
      setResendState("idle");
    }
  };

  if (isAuthLoading) {
    return (
      <AuthShell>
        <Box sx={{ display: "grid", placeItems: "center", py: 8 }}>
          <CircularProgress />
        </Box>
      </AuthShell>
    );
  }

  if (user && !forceForm) {
    return (
      <AuthShell subtitle="Account">
        <Seo title="Continue | Amarolio Account" />
        <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
          <Stack
            spacing={2.5}
            sx={{ alignItems: "center", textAlign: "center" }}
          >
            <Avatar
              sx={{
                width: 64,
                height: 64,
                fontWeight: 700,
                fontSize: "1.6rem",
                fontFamily: "'Poppins', 'Open Sans', sans-serif",
                background:
                  "linear-gradient(160deg, #8FB3F0 0%, #3D6BD4 100%)",
              }}
            >
              {user.username.charAt(0).toUpperCase()}
            </Avatar>
            <Box>
              <Typography variant="h6" sx={{ fontWeight: 700 }}>
                Continue as {user.username}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                You are already signed in to your Amarolio account.
              </Typography>
            </Box>
            <Button
              fullWidth
              variant="contained"
              href={getPersistedRedirectUri()}
              sx={{ py: 1.4 }}
            >
              Continue
            </Button>
            <Button
              fullWidth
              variant="text"
              onClick={() => setForceForm(true)}
            >
              Use another account
            </Button>
          </Stack>
        </Paper>
      </AuthShell>
    );
  }

  return (
    <AuthShell subtitle="Sign in">
      <Seo title="Sign in | Amarolio Account" />
      <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 } }}>
        {flow.step === "credentials" ? (
          <Box
            component="form"
            onSubmit={(event) => {
              event.preventDefault();
              flow.submitCredentials();
            }}
          >
            <Typography variant="h5" sx={{ fontWeight: 700, mb: 0.5 }}>
              Welcome back
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 3 }}>
              Sign in with your Amarolio account.
            </Typography>

            {flow.error && (
              <Alert severity="error" sx={{ mb: 2 }}>
                {flow.error}
              </Alert>
            )}

            {showUnverified && (
              <Box sx={{ mb: 2 }}>
                {resendState === "sent" ? (
                  <Alert severity="success">
                    Verification email sent. Check your inbox.
                  </Alert>
                ) : (
                  <Button
                    variant="outlined"
                    size="small"
                    fullWidth
                    startIcon={<MarkEmailReadOutlinedIcon />}
                    disabled={resendState === "sending" || !flow.email}
                    onClick={handleResendVerification}
                  >
                    {resendState === "sending"
                      ? "Sending..."
                      : "Resend verification email"}
                  </Button>
                )}
                {resendError && (
                  <Alert severity="error" sx={{ mt: 1 }}>
                    {resendError}
                  </Alert>
                )}
              </Box>
            )}

            <Stack spacing={2}>
              <TextField
                label="Email"
                type="email"
                fullWidth
                value={flow.email}
                onChange={(event) => flow.setEmail(event.target.value)}
                autoComplete="email"
              />
              <TextField
                label="Password"
                type="password"
                fullWidth
                value={flow.password}
                onChange={(event) => flow.setPassword(event.target.value)}
                autoComplete="current-password"
              />
            </Stack>

            <Button
              type="submit"
              fullWidth
              variant="contained"
              disabled={flow.isLoading}
              sx={{ mt: 3, py: 1.4 }}
            >
              {flow.isLoading ? "Signing in..." : "Sign in"}
            </Button>

            <Divider sx={{ my: 3 }}>
              <Typography variant="caption" color="text.secondary">
                New to Amarolio?
              </Typography>
            </Divider>

            <Button
              fullWidth
              variant="outlined"
              component={RouterLink}
              to={registerLink}
            >
              Create an account
            </Button>
          </Box>
        ) : (
          <Box>
            <Typography variant="h5" sx={{ fontWeight: 700, mb: 0.5 }}>
              Check your email
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 3 }}>
              {flow.info || "Enter the 6-digit code we sent you."}
            </Typography>

            {flow.error && (
              <Alert severity="error" sx={{ mb: 2 }}>
                {flow.error}
              </Alert>
            )}

            <OtpInput
              value={flow.otp}
              onChange={flow.setOtp}
              onComplete={(value) => flow.submitOtp(value)}
              disabled={flow.isLoading}
            />

            <Button
              fullWidth
              variant="contained"
              disabled={flow.isLoading}
              onClick={() => flow.submitOtp()}
              sx={{ mt: 3, py: 1.4 }}
            >
              {flow.isLoading ? "Verifying..." : "Verify and continue"}
            </Button>

            <Stack
              direction="row"
              sx={{ justifyContent: "space-between", alignItems: "center", mt: 1.5 }}
            >
              <Button
                size="small"
                startIcon={<ArrowBackIcon />}
                onClick={flow.backToCredentials}
              >
                Back
              </Button>
              <ResendButton
                seconds={flow.resendIn}
                onResend={flow.resend}
                disabled={flow.isLoading}
              />
            </Stack>
          </Box>
        )}
      </Paper>
    </AuthShell>
  );
}
