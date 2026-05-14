import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Box, CircularProgress, Typography, Container, Button } from '@mui/material'
import '../styles/RedirectPage.css'
import { Turnstile } from '@marsidev/react-turnstile'

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
    const [stars] = useState<FallingStar[]>(generateStars())

    const handleRedirect = async () => {
        try {
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

            const backendUrl = `${import.meta.env.VITE_SERVER_HOST || 'http://localhost:3000'}/api/v1/url/${slug}`

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

    useEffect(() => { }, [slug])
    return (
        <Box className="redirect-page">
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
                    <Typography variant="h6" component="div" className="brand-heading">
                        Amary
                        <Box className="leaf-icon">🍂</Box>
                    </Typography>

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
                                    variant="contained"
                                    onClick={handleRedirect}
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
