import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Box, CircularProgress, Typography, Container, Button } from '@mui/material'
import '../styles/RedirectPage.css'
import { Turnstile } from '@marsidev/react-turnstile'
import { Helmet } from 'react-helmet-async'
import type { MouseEvent } from 'react'

interface ServerResponse<T> {
    success: boolean
    data: T
}

interface URL {
    id: number
    user_id: string | null
    short_url: string
    url: string
    created_at: Date
    expired_at: Date | null
}

interface URLMetadata {
    url: URL
}

export interface ErrorResponse {
    detail: string;
}

interface FallingStar {
    id: number
    left: number
    delay: number
    duration: number
}

const generateStars = (): FallingStar[] =>
    Array.from({ length: 8 }, (_, i) => ({
        id: i,
        left: Math.random() * 100,
        delay: Math.random() * 0.5,
        duration: 2 + Math.random() * 1,
    }))

export default function RedirectPage() {
    const { slug } = useParams<{ slug: string }>()
    const [loading, setLoading] = useState(false)
    const [tsSuccess, setTsSuccess] = useState<boolean>(false)
    const [tsToken, setTsToken] = useState<string>("")
    const [error, setError] = useState<string | null>(null)
    const [destinationUrl, setDestinationUrl] = useState<string | null>(null)
    const [stars] = useState<FallingStar[]>(generateStars())

    const pageTitle = slug ? `Redirecting... | Amary` : "Amary | Secure Link Redirection";
    const mainSiteUrl = import.meta.env.VITE_AMARY_CLIENT_DOMAIN

    useEffect(() => {
        const fetchMetadata = async () => {
            if (!slug) return;
            try {
                const host = import.meta.env.VITE_SERVER_HOST
                const response = await fetch(`${host}/api/v1/url/${slug}/metadata`);
                if (!response.ok) {
                    const resBody: ServerResponse<ErrorResponse> = await response.json();
                    throw new Error(resBody.data.detail)
                }
                const resBody: ServerResponse<URLMetadata> = await response.json();
                if (resBody.data.url.url) {
                    setDestinationUrl(resBody.data.url.url);
                }
            } catch (err) {
                console.error("Failed to fetch link metadata for SEO", err);
            }
        };
        fetchMetadata();
    }, [slug]);

    const handleRedirect = async (e: MouseEvent<HTMLAnchorElement>) => {
        try {
            e.preventDefault()
            setLoading(true)
            setError(null)

            if (!slug) {
                setError('Invalid URL')
                setLoading(false)
                return
            }
            if (!tsToken) {
                setError("Unauthorized User")
                setLoading(false)
                return
            }
            await new Promise(resolve => setTimeout(resolve, 2000))

            const backendUrl = `${import.meta.env.VITE_SERVER_HOST}/api/v1/url/${slug}/redirect`

            window.location.href = `${backendUrl}?token=${encodeURIComponent(tsToken)}`
        } catch (err) {
            setError(err instanceof Error ? err.message : 'An error occurred')
            setLoading(false)
        }
    }

    const handleTsOnSuccess = (tk: string) => {
        setTsToken(tk)
        setTsSuccess(true)
    }

    return (
        <Box className="redirect-page">
            <Helmet>
                <title>{pageTitle}</title>
                <meta name="description" content="Please wait while Amary securely prepares your link. We are redirecting you to your destination." />
                <meta name="robots" content="noindex, follow" />
                {destinationUrl && <link rel="canonical" href={destinationUrl} />}
                <meta property="og:title" content={pageTitle} />
                <meta property="og:description" content="Secure and fast link redirection by Amary URL Shortener." />
                <meta property="og:type" content="website" />
                <script type="application/ld+json">
                    {JSON.stringify({
                        "@context": "https://schema.org",
                        "@type": "WebPage",
                        "name": pageTitle,
                        "isPartOf": {
                            "@type": "WebApplication",
                            "name": "Amary URL Shortener",
                            "url": mainSiteUrl,
                            "sameAs": destinationUrl ? [destinationUrl] : []
                        },
                        "description": "Secure intermediate redirection page for Amary links."
                    })}
                </script>
            </Helmet>
            {/* Falling leaves/stars background */}
            <div className="stars-container">
                {stars.map((star) => (
                    <div
                        key={star.id}
                        className="falling-star"
                        style={
                            {
                                '--left': `${star.left}%`,
                                '--delay': `${star.delay}s`,
                                '--duration': `${star.duration}s`,
                            } as React.CSSProperties
                        }
                    />
                ))}
            </div>

            <Container maxWidth="sm">
                <Box className="content-wrapper">
                    <a
                        href={mainSiteUrl}
                        style={{ textDecoration: 'none', color: 'inherit' }}
                        title="Go to Amary URL Shortener Home"
                    >
                        <Typography variant="h6" component="div" className="brand-heading">
                            Amary
                            <Box className="leaf-icon">🍂</Box>
                        </Typography>
                    </a>

                    <Typography variant="h4" component="h1" className="title">
                        Redirecting
                    </Typography>

                    <Typography variant="body1" className="subtitle">
                        {error ? 'An error occurred' : 'Please wait while we prepare your journey...'}
                    </Typography>

                    {error ? (
                        <Typography variant="body2" className="error-message">
                            {error}
                        </Typography>
                    ) : (
                        <>
                            <Box className="button-container">
                                <Button
                                    component="a"
                                    href={destinationUrl || '#'}
                                    variant="contained"
                                    onClick={handleRedirect}
                                    rel="follow"
                                    disabled={loading || !tsSuccess}
                                    sx={{
                                        marginTop: '2rem',
                                        backgroundColor: '#C87137',
                                        color: '#fff',
                                        fontWeight: 600,
                                        padding: '0.75rem 2rem',
                                        fontSize: '1rem',
                                        textTransform: 'none',
                                        borderRadius: '8px',
                                        display: 'flex',
                                        alignItems: 'center',
                                        gap: '0.5rem',
                                        '&:hover': {
                                            backgroundColor: '#b8631f',
                                        },
                                        '&:disabled': {
                                            backgroundColor: '#d4a574',
                                            color: '#fff',
                                        },
                                    }}
                                >
                                    {loading && (
                                        <CircularProgress
                                            size={20}
                                            sx={{
                                                color: '#fff',
                                            }}
                                        />
                                    )}
                                    {loading ? 'Redirecting...' : 'Continue'}
                                </Button>
                            </Box>
                        </>
                    )}

                    <Typography variant="caption" className="description">
                        {loading ? 'Preparing your journey...' : 'Click the button to proceed'}
                    </Typography>
                    <Turnstile onSuccess={handleTsOnSuccess} siteKey={import.meta.env.VITE_CF_TURNSTILE_SITEKEY} />
                </Box>
            </Container>
        </Box>
    )
}
