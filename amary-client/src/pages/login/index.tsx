import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Box, Container, Typography, Button, Alert, Paper } from '@mui/material';
import { Google as GoogleIcon } from '@mui/icons-material';
import { useAuth } from '../../controllers/useAuth';

export function LoginPage() {
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleGoogleLogin = async () => {
    setError('');
    setIsLoading(true);

    try {
      // Mock Google login - in production, this would use Google OAuth
      await login('google');
      navigate('/');
    } catch {
      setError('Google login failed. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ my: 8 }}>
        <Typography 
          variant="h3" 
          component="h1" 
          gutterBottom
          sx={{
            color: '#c25e00',
            fontWeight: 700,
            textAlign: 'center',
            mb: 2,
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
          sx={{ 
            p: 4, 
            border: '1px solid #e8dcc8',
            borderRadius: 2,
          }}
        >
          <Button
            fullWidth
            variant="outlined"
            onClick={handleGoogleLogin}
            disabled={isLoading}
            startIcon={<GoogleIcon />}
            sx={{
              py: 1.5,
              mb: 2,
              borderColor: '#faf6f0',
              color: '#faf6f0',
              fontWeight: 600,
              '&:hover': {
                borderColor: '#8b4513',
                backgroundColor: 'rgba(194, 94, 0, 0.04)',
                color: '#8b4513',
              },
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
