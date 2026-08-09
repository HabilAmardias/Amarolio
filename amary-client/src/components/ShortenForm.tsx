import { Box, TextField, Button, CircularProgress, Alert, Checkbox, FormControlLabel } from '@mui/material';
import { useShorten } from '../controllers/useShorten';
import { Turnstile, type TurnstileInstance } from '@marsidev/react-turnstile';
import { useRef } from 'react';
import { useAtom } from 'jotai';
import { userModel } from '../models/user/model';
import { ResultCard } from './ResultCard';

export function ShortenForm() {
  const ref = useRef<TurnstileInstance | null>(null)
  const [user] = useAtom(userModel.userAtom)
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
    result
  } = useShorten();

  const turnstileOnExpire = () => {
    ref.current.reset()
    setToken("")
  }

  return (
    <>
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
            border: '1px solid',
            borderColor: 'divider',
            borderRadius: 1,
            background: 'linear-gradient(145deg, rgba(255,255,255,0.6) 0%, rgba(255, 244, 230, 0.35) 100%)',
          }}>
            <Box sx={{
              display: 'flex',
              flexDirection: { xs: 'column', sm: 'row' },
              alignItems: { xs: 'flex-start', sm: 'center' },
              gap: 2,
              mb: 2,
            }}>
              <TextField
                type="number"
                label="Expiration (days)"
                value={expiresInDays || ''}
                onChange={(e) => setExpiresInDays(parseInt(e.target.value) || null)}
                disabled={noExpiry || isLoading}
                sx={{ width: { xs: '100%', sm: 180 } }}
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
            <Box sx={{ position: 'relative', mt: 2, overflow: 'visible' }}>
              <TextField
                fullWidth
                label="Custom slug (optional)"
                placeholder="e.g. my-special-link"
                value={rawCustom}
                onChange={(e) => onCustomChange(e.target.value)}
                disabled={isLoading}
                helperText={isCheckingSlug ? 'Checking availability...' : 'Leave empty to generate automatically'}
                sx={{
                  '& .MuiInputBase-input': {
                    paddingRight: '2.5rem',
                  },
                }}
              />

              {isCheckingSlug && (
                <Box sx={{ position: 'absolute', right: 12, top: 50, bottom: 0, display: { xs: 'none', sm: 'flex' }, alignItems: 'center', pointerEvents: 'none' }}>
                  <CircularProgress size={16} sx={{ color: 'primary.main' }} />
                </Box>
              )}

              {isCheckingSlug && (
                <Box sx={{ display: { xs: 'block', sm: 'none' }, mt: 1, textAlign: 'center' }}>
                  <CircularProgress size={12} sx={{ color: 'primary.main' }} />
                </Box>
              )}
            </Box>
          </Box>
        )}
        <Button
          sx={{
            mb: 2,
            py: 1.2,
          }}
          fullWidth
          variant="contained"
          onClick={handleShorten}
          disabled={isLoading || !url || !token || isCheckingSlug}
        >
          {isLoading ? <CircularProgress size={24} color="inherit" /> : 'Shorten URL'}
        </Button>

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

        <Box sx={{ display: 'flex', justifyContent: 'center', mt: 1 }}>
          <Turnstile ref={ref} onExpire={turnstileOnExpire} onSuccess={(tk) => setToken(tk)} siteKey={import.meta.env.VITE_CF_TURNSTILE_SITEKEY} />
        </Box>
      </Box>

      {result && <ResultCard result={result} />}
    </>
  );
}
