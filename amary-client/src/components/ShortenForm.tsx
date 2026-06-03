import { Box, TextField, Button, CircularProgress, Alert, Checkbox, FormControlLabel } from '@mui/material';
import { useShorten } from '../controllers/useShorten';
import { useAuth } from '../controllers/useAuth';
import { Turnstile, type TurnstileInstance } from '@marsidev/react-turnstile';
import { useRef } from 'react';

export function ShortenForm() {
  const ref = useRef<TurnstileInstance | null>(null)
  const { user } = useAuth();
  const {
    url,
    setUrl,
    expiresInDays,
    setExpiresInDays,
    noExpiry,
    setNoExpiry,
    token,
    setToken,
    error,
    isLoading,
    handleShorten,
  } = useShorten();


  const turnstileOnExpire = () => {
    ref.current.reset()
    setToken("")
  }

  return (
    <Box sx={{ mb: 3 }}>
      <TextField
        fullWidth
        label="Enter URL to shorten"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        disabled={isLoading}
        sx={{ mb: 2 }}
      />

      {user && (
        <Box sx={{
          mb: 2,
          p: 2,
          border: '1px solid #e8dcc8',
          borderRadius: 2,
          background: '#faf6f0',
        }}>
          <TextField
            type="number"
            label="Expiration (days)"
            value={expiresInDays || ''}
            onChange={(e) => setExpiresInDays(parseInt(e.target.value) || null)}
            disabled={noExpiry || isLoading}
            sx={{ mr: 2 }}
          />
          <FormControlLabel
            control={
              <Checkbox
                checked={noExpiry}
                onChange={(e) => setNoExpiry(e.target.checked)}
                disabled={isLoading}
              />
            }
            label="No expiration"
          />
        </Box>
      )}
      <Button
        sx={{
          mb: 2
        }}
        fullWidth
        variant="contained"
        onClick={handleShorten}
        disabled={isLoading || !url || !token}
      >
        {isLoading ? <CircularProgress size={24} /> : 'Shorten URL'}
      </Button>

      <Turnstile ref={ref} onExpire={turnstileOnExpire} onSuccess={(tk) => setToken(tk)} siteKey={import.meta.env.VITE_CF_TURNSTILE_SITEKEY} />

      {error && (
        <Alert
          severity="error"
          sx={{
            mt: 2,
          }}
        >
          {error}
        </Alert>
      )}
    </Box>
  );
}
