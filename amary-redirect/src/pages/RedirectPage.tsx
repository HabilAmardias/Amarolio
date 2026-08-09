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

interface FallingLeaf {
    id: number
    left: number
    size: number
    delay: number
    duration: number
    drift: number
    spin: number
    sway: number
    color: string
}

const LEAF_COLORS = ['#d97706', '#ea580c', '#f59e0b', '#b45309', '#c2410c']

const generateLeaves = (): FallingLeaf[] =>
    Array.from({ length: 12 }, (_, i) => ({
        id: i,
        left: Math.random() * 100,
        size: 14 + Math.random() * 16,
        delay: Math.random() * 2.5,
        duration: 9 + Math.random() * 6,
        drift: (Math.random() - 0.5) * 220,
        spin: 180 + Math.random() * 360,
        sway: 2 + Math.random() * 2.5,
        color: LEAF_COLORS[i % LEAF_COLORS.length],
    }))

function LeafIcon({ variant = 'filled' }: { variant?: 'filled' | 'outline' }) {
    const shared = {
        viewBox: '0 0 24 24',
        'aria-hidden': true,
        focusable: false,
    }
    if (variant === 'filled') {
        return (
            <svg {...shared} fill="currentColor">
                <path d="M11 20A7 7 0 0 1 9.8 6.1C15.5 5 17 4.48 19 2c1 2 2 4.18 2 8 0 5.5-4.78 10-10 10Z" />
                <path d="M2 21c0-3 1.85-5.36 5.08-6C9.5 14.52 12 13 13 12" fill="rgba(255,255,255,0.3)" />
            </svg>
        )
    }
    return (
        <svg {...shared} fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M11 20A7 7 0 0 1 9.8 6.1C15.5 5 17 4.48 19 2c1 2 2 4.18 2 8 0 5.5-4.78 10-10 10Z" />
            <path d="M2 21c0-3 1.85-5.36 5.08-6C9.5 14.52 12 13 13 12" />
        </svg>
    )
}

export default function RedirectPage() {
    const { slug } = useParams<{ slug: string }>()
    const [loading, setLoading] = useState(false)
    const [tsSuccess, setTsSuccess] = useState<boolean>(false)
    const [tsToken, setTsToken] = useState<string>("")
    const [error, setError] = useState<string | null>(null)
    const [destinationUrl, setDestinationUrl] = useState<string | null>(null)
    const [leaves] = useState<FallingLeaf[]>(generateLeaves())

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
                setError(err instanceof Error ? err.message : 'An error occurred')
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
                <meta property="og:site_name" content="Amary" />
                <meta property="og:title" content={pageTitle} />
                <meta property="og:description" content="Secure and fast link redirection by Amary URL Shortener." />
                <meta property="og:type" content="website" />
                <meta property="og:url" content={window.location.href} />
                <meta property="og:image" content="https://amary.id/og-image.png" />
                <meta property="og:image:width" content="1200" />
                <meta property="og:image:height" content="630" />
                <meta property="og:image:alt" content="Amary URL shortener" />
                <meta name="twitter:card" content="summary_large_image" />
                <meta name="twitter:title" content={pageTitle} />
                <meta name="twitter:description" content="Secure and fast link redirection by Amary URL Shortener." />
                <meta name="twitter:image" content="https://amary.id/og-image.png" />
                <meta name="twitter:image:alt" content="Amary URL shortener" />
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
            {/* Ambient autumn light behind the glass */}
            <div className="ambient-orbs" aria-hidden="true">
                <span className="orb orb-1" />
                <span className="orb orb-2" />
                <span className="orb orb-3" />
                <span className="orb orb-4" />
            </div>

            {/* Falling leaves background */}
            <div className="leaves-container" aria-hidden="true">
                {leaves.map((leaf) => (
                    <span
                        key={leaf.id}
                        className="leaf-item"
                        style={
                            {
                                '--left': `${leaf.left}%`,
                                '--size': `${leaf.size}px`,
                                '--delay': `${leaf.delay}s`,
                                '--duration': `${leaf.duration}s`,
                                '--drift': `${leaf.drift}px`,
                                '--spin': `${leaf.spin}deg`,
                                '--sway': `${leaf.sway}s`,
                                '--leaf-color': leaf.color,
                            } as React.CSSProperties
                        }
                    >
                        <LeafIcon />
                    </span>
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
                            <Box component="span" className="leaf-icon">
                                <LeafIcon variant="outline" />
                            </Box>
                        </Typography>
                    </a>

                    <Typography variant="h4" component="h1" className="title">
                        Redirecting
                    </Typography>

                    <Typography variant="body1" className="subtitle">
                        {error ? 'An error occurred' : 'Please wait while we prepare your journey...'}
                    </Typography>

                    {error ? (
                        <Typography sx={{
                            marginY: "1rem",
                        }} variant="body2" className="error-message">
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
                                        background: 'linear-gradient(135deg, #ee8a3c 0%, #c87137 55%, #b45309 100%)',
                                        color: '#fff',
                                        fontFamily: "'Inter', sans-serif",
                                        fontWeight: 600,
                                        letterSpacing: '0.5px',
                                        padding: '0.9rem 2.75rem',
                                        fontSize: '1rem',
                                        textTransform: 'none',
                                        borderRadius: '14px',
                                        display: 'flex',
                                        alignItems: 'center',
                                        gap: '0.6rem',
                                        boxShadow: '0 12px 26px -10px rgba(180, 83, 9, 0.6), inset 0 1px 0 rgba(255,255,255,0.25)',
                                        transition: 'transform 0.2s ease, box-shadow 0.2s ease, background 0.2s ease',
                                        '&:hover': {
                                            background: 'linear-gradient(135deg, #f0943f 0%, #b8631f 55%, #a3480a 100%)',
                                            boxShadow: '0 16px 30px -10px rgba(180, 83, 9, 0.65)',
                                            transform: 'translateY(-2px)',
                                        },
                                        '&:disabled': {
                                            background: 'rgba(217, 119, 6, 0.35)',
                                            color: 'rgba(255,255,255,0.95)',
                                            boxShadow: 'none',
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
                    <Box className="turnstile-wrapper">
                        <Turnstile onSuccess={handleTsOnSuccess} siteKey={import.meta.env.VITE_CF_TURNSTILE_SITEKEY} />
                    </Box>
                </Box>
            </Container>
        </Box>
    )
}
