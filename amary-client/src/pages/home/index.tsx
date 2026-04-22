import { Box, Container, Typography } from '@mui/material';
import { ShortenForm } from '../../components/ShortenForm';
import { ResultCard } from '../../components/ResultCard';
import { useShorten } from '../../controllers/useShorten';

export function HomePage() {
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

        <ShortenForm />

        {result && <ResultCard result={result} />}
      </Box>
    </Container>
  );
}
