const EMAIL_RE = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
const PASSWORD_RE = /^[a-zA-Z0-9]{8,13}$/;

export function validateEmail(email: string): string | null {
  if (!email.trim()) return "Email is required.";
  if (!EMAIL_RE.test(email.trim())) return "Enter a valid email address.";
  return null;
}

export function validatePassword(password: string): string | null {
  if (!password) return "Password is required.";
  if (!PASSWORD_RE.test(password)) {
    return "Password must be 8–13 characters, letters and numbers only.";
  }
  return null;
}

export function validateOtp(otp: string): string | null {
  if (!/^\d{6}$/.test(otp)) return "Enter the 6-digit code from your email.";
  return null;
}
