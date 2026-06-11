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
    rawCustom,
    onCustomChange,
    isCheckingSlug,
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
          <Box sx={{ position: 'relative', mt: 2 }}>
            <TextField
              fullWidth
              label="Custom slug (optional)"
              placeholder="e.g. my-special-link"
              value={rawCustom}
              onChange={(e) => onCustomChange(e.target.value)}
              disabled={isLoading}
              helperText={isCheckingSlug ? 'Checking availability...' : 'Leave empty to generate automatically'}
            />

            {/* desktop: small spinner inside the input on the right */}
            {isCheckingSlug && (
              <Box sx={{ position: 'absolute', right: 12, top: '50%', transform: 'translateY(-50%)', display: { xs: 'none', sm: 'flex' }, alignItems: 'center' }}>
                <CircularProgress size={16} />
              </Box>
            )}

            {/* mobile: show a smaller spinner under the field */}
            {isCheckingSlug && (
              <Box sx={{ display: { xs: 'block', sm: 'none' }, mt: 1 }}>
                <CircularProgress size={12} />
              </Box>
            )}
          </Box>
        </Box>
      )}
      <Button
        sx={{
          mb: 2
        }}
        fullWidth
        variant="contained"
        onClick={handleShorten}
        disabled={isLoading || !url || !token || isCheckingSlug}
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
