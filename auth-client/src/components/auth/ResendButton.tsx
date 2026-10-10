import { Button } from "@mui/material";

interface ResendButtonProps {
  seconds: number;
  onResend: () => void;
  disabled?: boolean;
}

export function ResendButton({
  seconds,
  onResend,
  disabled,
}: ResendButtonProps) {
  return (
    <Button
      onClick={onResend}
      disabled={seconds > 0 || disabled}
      size="small"
      sx={{ fontWeight: 600 }}
    >
      {seconds > 0 ? `Resend code in ${seconds}s` : "Resend code"}
    </Button>
  );
}
