import { Card, CardContent, Typography, Box, IconButton, Snackbar } from '@mui/material';
import { ContentCopy } from '@mui/icons-material';
import { useState } from 'react';
import type { ShortenResponse } from '../models/url.model';

interface ResultCardProps {
  result: ShortenResponse;
}

export function ResultCard({ result }: ResultCardProps) {
  const [open, setOpen] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(result.shortUrl);
    setOpen(true);
  };

  return (
    <>
      <Card variant="outlined">
        <CardContent>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
            <Typography 
              variant="h6" 
              sx={{ 
                flexGrow: 1,
                color: '#c25e00',
                fontWeight: 600,
              }}
            >
              {result.shortUrl}
            </Typography>
            <IconButton 
              onClick={handleCopy} 
              size="small"
              sx={{
                color: '#c25e00',
                '&:hover': {
                  backgroundColor: 'rgba(194, 94, 0, 0.08)',
                },
              }}
            >
              <ContentCopy />
            </IconButton>
          </Box>
          <Typography 
            variant="body2" 
            sx={{ 
              mb: 1,
              color: 'text.secondary',
            }}
          >
            Original: {result.originalUrl}
          </Typography>
          <Typography 
            variant="caption" 
            sx={{
              color: '#8b4513',
              fontWeight: 600,
            }}
          >
            {result.expiresAt
              ? `Expires: ${new Date(result.expiresAt).toLocaleDateString()}`
              : 'No expiration'}
          </Typography>
        </CardContent>
      </Card>
      <Snackbar
        open={open}
        autoHideDuration={2000}
        onClose={() => setOpen(false)}
        message="Copied to clipboard"
      />
    </>
  );
}
