import { Box, TextField, Button, CircularProgress, Alert, Checkbox, FormControlLabel } from '@mui/material';
import { useShorten } from '../controllers/useShorten';
import { useAuth } from '../controllers/useAuth';

export function ShortenForm() {
  const { user } = useAuth();
  const {
    url,
    setUrl,
    expiresInDays,
    setExpiresInDays,
    noExpiry,
    setNoExpiry,
    error,
    isLoading,
    handleShorten,
  } = useShorten();

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
            inputProps={{ min: 1 }}
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
        fullWidth
        variant="contained"
        onClick={handleShorten}
        disabled={isLoading || !url}
      >
        {isLoading ? <CircularProgress size={24} /> : 'Shorten URL'}
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
    </Box>
  );
}
