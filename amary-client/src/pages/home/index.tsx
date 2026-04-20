import { Box, Container, Typography, Link as MuiLink } from '@mui/material';
import { ShortenForm } from '../../components/ShortenForm';
import { ResultCard } from '../../components/ResultCard';
import { useShorten } from '../../controllers/useShorten';
import { useAuth } from '../../controllers/useAuth';
import { Link } from 'react-router-dom';

export function HomePage() {
  const { user } = useAuth();
  const { result } = useShorten();

  return (
    <Container maxWidth="sm">
      <Box sx={{ my: 8, textAlign: 'center' }}>
        <Typography 
          variant="h2" 
          component="h1" 
          gutterBottom
          sx={{
            color: '#c25e00',
            fontWeight: 700,
            mb: 3,
          }}
        >
          Amary
        </Typography>
        <Typography 
          variant="body1" 
          color="text.secondary" 
          sx={{ 
            mb: 4,
            fontSize: '1.1rem',
          }}
        >
          Simplify your links with ease
        </Typography>

        {user && (
          <MuiLink 
            component={Link} 
            to="/dashboard" 
            sx={{ 
              mb: 4, 
              display: 'block',
              color: '#c25e00',
              textDecoration: 'none',
              fontWeight: 600,
              '&:hover': {
                textDecoration: 'underline',
              },
            }}
          >
            → Go to Dashboard
          </MuiLink>
        )}

        <ShortenForm />

        {result && <ResultCard result={result} />}
      </Box>
    </Container>
  );
}
