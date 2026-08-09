import { useState } from 'react';
import { Box, Container, Typography, Button, Alert, Paper } from '@mui/material';
import { Google as GoogleIcon } from '@mui/icons-material';
import { Navigate } from 'react-router-dom';
import { Helmet } from 'react-helmet-async';
import { useAtom } from 'jotai';
import { userModel } from '../../models/user/model';
import { login } from '../../api/auth.api';

export function LoginPage() {
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [user] = useAtom(userModel.userAtom);

  const handleGoogleLogin = async () => {
    setError('');
    setIsLoading(true);

    try {
      await login();
    } catch {
      setError('Google login failed. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  if (user) {
    return <Navigate to="/" replace />
  }

  return (
    <Container maxWidth="sm">
      <Helmet>
        <title>Sign In | Amary</title>
        <meta name="description" content="Sign in to your Amary account to manage your shortened links, track analytics, and create custom short URLs." />
        <meta name="robots" content="noindex, nofollow" />
        <meta name="googlebot" content="noindex, nofollow" />
      </Helmet>
      <Box sx={{ py: { xs: 6, sm: 8 }, display: 'flex', flexDirection: 'column', justifyContent: 'center', minHeight: 'calc(100vh - 64px)' }}>
        <Typography
          variant="h3"
          component="h1"
          gutterBottom
          sx={{
            color: 'primary.main',
            fontWeight: 700,
            textAlign: 'center',
            mb: 2,
            fontSize: { xs: '1.9rem', sm: '2.5rem' },
          }}
        >
          Welcome Back
        </Typography>

        <Typography
          variant="body1"
          color="text.secondary"
          sx={{
            textAlign: 'center',
            mb: 4,
          }}
        >
          Sign in with your Google account to continue
        </Typography>

        <Paper
          elevation={0}
          variant="outlined"
          sx={{
            p: { xs: 2.5, sm: 4 },
            borderRadius: 4,
          }}
        >
          <Button
            fullWidth
            variant="contained"
            onClick={() => handleGoogleLogin()}
            disabled={isLoading}
            startIcon={<GoogleIcon />}
            sx={{
              py: 1.5,
            }}
          >
            {isLoading ? 'Signing in...' : 'Sign in with Google'}
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
        </Paper>

        <Typography
          variant="caption"
          color="text.secondary"
          sx={{
            display: 'block',
            textAlign: 'center',
            mt: 3,
          }}
        >
          By signing in, you agree to our Terms of Service and Privacy Policy
        </Typography>
      </Box>
    </Container>
  );
}
