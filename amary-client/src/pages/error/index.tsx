import { Box, Container, Typography, Button } from '@mui/material';
import { Link } from 'react-router-dom';
import { Helmet } from 'react-helmet-async';
import { HelpOutlined } from '@mui/icons-material';

export function Error404Page() {
    return (
        <Container maxWidth="sm">
            <Helmet>
                <title>404 - Page Not Found | Amary</title>
                <meta name="description" content="The page you are looking for does not exist or has been moved." />
            </Helmet>
            <Box
                sx={{
                    py: 8,
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'center',
                    alignItems: 'center',
                    textAlign: 'center',
                    minHeight: 'calc(100vh - 120px)',
                }}
            >
                <HelpOutlined
                    sx={{
                        fontSize: '6rem',
                        color: 'primary.main',
                        mb: 3,
                        animation: 'float 3s ease-in-out infinite',
                        '@keyframes float': {
                            '0%, 100%': {
                                transform: 'translateY(0)',
                            },
                            '50%': {
                                transform: 'translateY(-10px)',
                            },
                        },
                    }}
                />
                <Typography
                    variant="h2"
                    component="h1"
                    gutterBottom
                    sx={{
                        color: 'primary.main',
                        fontWeight: 700,
                        mb: 2,
                    }}
                >
                    404
                </Typography>
                <Typography
                    variant="h5"
                    component="h2"
                    gutterBottom
                    sx={{
                        color: 'text.primary',
                        mb: 2,
                    }}
                >
                    Page Not Found
                </Typography>
                <Typography
                    variant="body1"
                    color="text.secondary"
                    sx={{
                        mb: 4,
                        fontSize: '1.1rem',
                        maxWidth: '400px',
                    }}
                >
                    The link you followed might be broken, or the page may have been removed. Let's get you back on track.
                </Typography>
                <Button
                    component={Link}
                    to="/"
                    variant="contained"
                    size="large"
                    sx={{
                        px: 4,
                        py: 1.5,
                        fontSize: '1rem',
                    }}
                >
                    Go Back Home
                </Button>
            </Box>
        </Container>
    );
}
