import { Box, Container, Typography, Paper } from '@mui/material';
import { ShortenForm } from '../../components/ShortenForm';
import { Helmet } from 'react-helmet-async';
import { Link } from 'react-router-dom';
import { useAuth } from '../../controllers/useAuth';

export function HomePage() {
  const { user } = useAuth();

  return (
    <Container maxWidth="sm">
      <Helmet>
        <title>Amary | Fast & Simple URL Shortener</title>
        <meta name="description" content="Amary helps you simplify and shorten long URLs into manageable links. Free, fast, and easy to use for all your sharing needs." />
        <meta name="keywords" content="URL shortener, link shortener, Amary, simplify links, marketing tools" />
        <link rel="canonical" href={window.location.origin} />
        <meta property="og:site_name" content="Amary" />
        <meta property="og:title" content="Amary | Fast & Simple URL Shortener" />
        <meta property="og:description" content="Simplify your links with ease using Amary. Create short, manageable links in seconds." />
        <meta property="og:url" content={window.location.origin} />
        <meta property="og:type" content="website" />
        <meta property="og:image" content={`${window.location.origin}/og-image.png`} />
        <meta property="og:image:width" content="1200" />
        <meta property="og:image:height" content="630" />
        <meta property="og:image:alt" content="Amary - Fast and simple URL shortener" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content="Amary | Fast & Simple URL Shortener" />
        <meta name="twitter:description" content="Simplify your links with ease using Amary. Create short, manageable links in seconds." />
        <meta name="twitter:image" content={`${window.location.origin}/og-image.png`} />
        <meta name="twitter:image:alt" content="Amary - Fast and simple URL shortener" />
        <script type="application/ld+json">
          {JSON.stringify({
            "@context": "https://schema.org",
            "@type": "WebApplication",
            "name": "Amary",
            "url": window.location.origin,
            "description": "Simplify your links with ease using Amary URL Shortener.",
            "applicationCategory": "UtilitiesApplication",
            "offers": {
              "@type": "Offer",
              "price": "0",
              "priceCurrency": "USD"
            },
            "featureList": [
              "Free URL shortening",
              "Custom expiration",
              "Click analytics",
              "No signup required"
            ],
            "publisher": {
              "@type": "Organization",
              "name": "Amary",
              "logo": {
                "@type": "ImageObject",
                "url": `${window.location.origin}/favicon.svg`
              }
            }
          })}
        </script>
        <script type="application/ld+json">
          {JSON.stringify({
            "@context": "https://schema.org",
            "@type": "WebSite",
            "name": "Amary",
            "url": window.location.origin,
            "potentialAction": {
              "@type": "SearchAction",
              "target": `${window.location.origin}/?query={search_term_string}`,
              "query-input": "required name=search_term_string"
            }
          })}
        </script>
      </Helmet>
      <Box sx={{ py: { xs: 6, sm: 8 }, display: 'flex', flexDirection: 'column', justifyContent: 'center', textAlign: 'center', minHeight: 'calc(100vh - 64px)' }}>
        <Typography
          variant="h2"
          component="h1"
          gutterBottom
          sx={{
            color: 'primary.main',
            fontWeight: 700,
            mb: 2,
            fontSize: { xs: '2rem', sm: '2.8rem' },
          }}
        >
          Amary
        </Typography>
        <Typography
          variant="body1"
          color="text.secondary"
          sx={{
            mb: 4,
            fontSize: { xs: '1rem', sm: '1.1rem' },
          }}
        >
          Chill out, we've got your links covered — nice and short.
        </Typography>

        <Paper
          elevation={0}
          variant="outlined"
          sx={{
            p: { xs: 2, sm: 3.5 },
            mb: 3,
          }}
        >
          <ShortenForm />
        </Paper>

        <Box
          sx={{
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'center',
            gap: { xs: 1.5, sm: 2 },
            mb: 3,
          }}
        >
          {['Free', 'Fast', 'No signup needed'].map((feature) => (
            <Box
              key={feature}
              component="span"
              sx={{
                px: 1.5,
                py: 0.75,
                borderRadius: '999px',
                background: 'linear-gradient(145deg, rgba(255,255,255,0.7) 0%, rgba(255, 244, 230, 0.45) 100%)',
                backdropFilter: 'blur(12px) saturate(150%)',
                WebkitBackdropFilter: 'blur(12px) saturate(150%)',
                border: '1px solid',
                borderColor: 'rgba(255, 255, 255, 0.7)',
                boxShadow: 'inset 0 1px 0 rgba(255, 255, 255, 0.7), 0 4px 12px rgba(120, 53, 15, 0.1)',
                color: 'text.secondary',
                fontSize: '0.9rem',
                fontWeight: 600,
              }}
            >
              {feature}
            </Box>
          ))}
        </Box>

        {user && (
          <Box sx={{ mt: 1 }}>
            <Typography
              component={Link}
              to="/dashboard"
              variant="body2"
              sx={{
                color: 'primary.main',
                textDecoration: 'none',
                '&:hover': { textDecoration: 'underline' },
              }}
            >
              Go to your dashboard →
            </Typography>
          </Box>
        )}
      </Box>
    </Container>
  );
}
