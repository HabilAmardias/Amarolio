import { useCallback, useEffect, useRef, useState } from "react";
import { login, prelogin, resendOtp } from "../api/auth.api";
import { getPersistedRedirectUri } from "../lib/redirect";
import { toApiError } from "../lib/errors";
import {
  validateEmail,
  validateOtp,
  validatePassword,
} from "../lib/validation";

const OTP_COOLDOWN_SECONDS = 60;

type Step = "credentials" | "otp";

export function useLoginFlow(initialEmail = "") {
  const [step, setStep] = useState<Step>("credentials");
  const [email, setEmail] = useState(initialEmail);
  const [password, setPassword] = useState("");
  const [otp, setOtp] = useState("");
  const [error, setError] = useState("");
  const [errorCode, setErrorCode] = useState<number | undefined>(undefined);
  const [info, setInfo] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [resendIn, setResendIn] = useState(0);
  const timer = useRef<number | null>(null);

  useEffect(() => {
    if (resendIn <= 0) return;
    timer.current = window.setInterval(() => {
      setResendIn((value) => (value <= 1 ? 0 : value - 1));
    }, 1000);
    return () => {
      if (timer.current) window.clearInterval(timer.current);
    };
  }, [resendIn]);

  const submitCredentials = useCallback(async () => {
    const emailError = validateEmail(email);
    const passwordError = validatePassword(password);
    if (emailError || passwordError) {
      setError(emailError ?? passwordError ?? "");
      return;
    }

    setError("");
    setErrorCode(undefined);
    setInfo("");
    setIsLoading(true);
    try {
      await prelogin(email.trim(), password);
      setStep("otp");
      setResendIn(OTP_COOLDOWN_SECONDS);
      setInfo(`We emailed a 6-digit code to ${email.trim()}.`);
    } catch (err) {
      const apiError = toApiError(err);
      setError(apiError.message);
      setErrorCode(apiError.code);
    } finally {
      setIsLoading(false);
    }
  }, [email, password]);

  const submitOtp = useCallback(
    async (value?: string) => {
      const code = value ?? otp;
      const otpError = validateOtp(code);
      if (otpError) {
        setError(otpError);
        return;
      }

      setError("");
      setIsLoading(true);
      try {
        await login(code);
        window.location.href = getPersistedRedirectUri();
      } catch (err) {
        const apiError = toApiError(err);
        setError(apiError.message);
        setErrorCode(apiError.code);
        setIsLoading(false);
      }
    },
    [otp]
  );

  const resend = useCallback(async () => {
    if (resendIn > 0) return;
    setError("");
    setErrorCode(undefined);
    try {
      await resendOtp();
      setResendIn(OTP_COOLDOWN_SECONDS);
      setInfo("A new code is on its way.");
    } catch (err) {
      setError(toApiError(err).message);
    }
  }, [resendIn]);

  const backToCredentials = useCallback(() => {
    setStep("credentials");
    setOtp("");
    setError("");
    setErrorCode(undefined);
    setInfo("");
  }, []);

  return {
    step,
    email,
    setEmail,
    password,
    setPassword,
    otp,
    setOtp,
    error,
    errorCode,
    info,
    isLoading,
    resendIn,
    submitCredentials,
    submitOtp,
    resend,
    backToCredentials,
  };
}
