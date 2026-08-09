import { Card, CardContent, Typography, Box, IconButton, Snackbar } from '@mui/material';
import { ContentCopy } from '@mui/icons-material';
import { useState } from 'react';
import type { ShortenResponse } from '../models/url/type';

interface ResultCardProps {
  result: ShortenResponse;
}

export function ResultCard({ result }: ResultCardProps) {
  const [open, setOpen] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(result.url);
    setOpen(true);
  };

  return (
    <>
      <Card variant="outlined">
        <CardContent sx={{ p: { xs: 2, sm: 3 } }}>
          <Box sx={{ display: 'flex', alignItems: 'center', mb: 1, gap: 1 }}>
            <Typography
              variant="h6"
              sx={{
                flexGrow: 1,
                color: 'primary.main',
                fontWeight: 600,
                wordBreak: 'break-all',
                fontSize: { xs: '1rem', sm: '1.25rem' },
              }}
            >
              {result.url}
            </Typography>
            <IconButton
              onClick={handleCopy}
              size="small"
              aria-label="Copy short URL"
              sx={{
                color: 'primary.main',
                flexShrink: 0,
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
              wordBreak: 'break-all',
            }}
          >
            Original: {result.original_url}
          </Typography>
          <Typography
            variant="caption"
            sx={{
              color: 'secondary.main',
              fontWeight: 600,
            }}
          >
            {result.expired_at
              ? `Expires: ${new Date(result.expired_at).toLocaleDateString()}`
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
