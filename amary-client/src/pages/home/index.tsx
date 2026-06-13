import { Box, Container, Typography } from '@mui/material';
import { ShortenForm } from '../../components/ShortenForm';
import { ResultCard } from '../../components/ResultCard';
import { useShorten } from '../../controllers/useShorten';
import { Helmet } from 'react-helmet-async';

export function HomePage() {
  const { result } = useShorten();

  return (
    <Container maxWidth="sm">
      <Helmet>
        <title>Amary | Fast & Simple URL Shortener</title>
        <meta name="description" content="Amary helps you simplify and shorten long URLs into manageable links. Free, fast, and easy to use for all your sharing needs." />
        <meta name="keywords" content="URL shortener, link shortener, Amary, simplify links, marketing tools" />
        <link rel="canonical" href={window.location.origin} />
        <meta property="og:title" content="Amary | Fast & Simple URL Shortener" />
        <meta property="og:description" content="Simplify your links with ease using Amary. Create short, manageable links in seconds." />
        <meta property="og:url" content={window.location.origin} />
        <meta property="og:type" content="website" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content="Amary | URL Shortener" />
        <script type="application/ld+json">
          {JSON.stringify({
            "@context": "https://schema.org",
            "@type": "WebApplication",
            "name": "Amary",
            "url": window.location.origin,
            "description": "Simplify your links with ease using Amary URL Shortener.",
            "applicationCategory": "UtilitiesApplication"
          })}
        </script>
      </Helmet>
      <Box sx={{ py: 8, display: 'flex', flexDirection: 'column', justifyContent: 'center', textAlign: 'center', minHeight: 'calc(100vh - 64px)' }}>
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
