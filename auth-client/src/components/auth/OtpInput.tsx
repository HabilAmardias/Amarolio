import { useRef } from "react";
import { Box, TextField } from "@mui/material";

const LENGTH = 6;

interface OtpInputProps {
  value: string;
  onChange: (value: string) => void;
  onComplete?: (value: string) => void;
  disabled?: boolean;
}

export function OtpInput({
  value,
  onChange,
  onComplete,
  disabled,
}: OtpInputProps) {
  const refs = useRef<Array<HTMLInputElement | null>>([]);
  const digits = Array.from(
    { length: LENGTH },
    (_, index) => value[index] ?? ""
  );

  const commit = (next: string[]) => {
    const joined = next.join("").slice(0, LENGTH);
    onChange(joined);
    if (joined.length === LENGTH) onComplete?.(joined);
  };

  const handleChange = (index: number, raw: string) => {
    const cleaned = raw.replace(/\D/g, "");
    const next = [...digits];

    if (!cleaned) {
      next[index] = "";
      commit(next);
      return;
    }

    for (let i = 0; i < cleaned.length && index + i < LENGTH; i += 1) {
      next[index + i] = cleaned[i];
    }
    commit(next);

    const focusIndex = Math.min(index + cleaned.length, LENGTH - 1);
    refs.current[focusIndex]?.focus();
  };

  const handleKeyDown = (index: number, event: React.KeyboardEvent) => {
    if (event.key === "Backspace" && !digits[index] && index > 0) {
      event.preventDefault();
      const next = [...digits];
      next[index - 1] = "";
      commit(next);
      refs.current[index - 1]?.focus();
    }
    if (event.key === "ArrowLeft" && index > 0) {
      refs.current[index - 1]?.focus();
    }
    if (event.key === "ArrowRight" && index < LENGTH - 1) {
      refs.current[index + 1]?.focus();
    }
  };

  return (
    <Box
      sx={{
        display: "flex",
        gap: { xs: 0.75, sm: 1 },
        justifyContent: "space-between",
      }}
    >
      {digits.map((digit, index) => (
        <TextField
          key={index}
          inputRef={(el) => {
            refs.current[index] = el;
          }}
          value={digit}
          onChange={(event) => handleChange(index, event.target.value)}
          onKeyDown={(event) => handleKeyDown(index, event)}
          onFocus={(event) => event.target.select()}
          disabled={disabled}
          autoFocus={index === 0}
          slotProps={{
            htmlInput: {
              inputMode: "numeric",
              pattern: "[0-9]*",
              autoComplete: "one-time-code",
              "aria-label": `Digit ${index + 1}`,
              style: {
                textAlign: "center",
                fontSize: "1.35rem",
                fontWeight: 600,
                padding: "10px 0",
              },
            },
          }}
          sx={{ flex: 1, minWidth: 0 }}
        />
      ))}
    </Box>
  );
}
