import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Box, CircularProgress, Typography, Container } from '@mui/material'
import '../styles/RedirectPage.css'

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
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)
    const [stars] = useState<FallingStar[]>(generateStars())

    useEffect(() => {
        const handleRedirect = async () => {
            try {
                if (!slug) {
                    setError('No redirect slug provided')
                    setLoading(false)
                    return
                }

                // Show loading animation for 2 seconds for better UX
                await new Promise(resolve => setTimeout(resolve, 2000))

                // Redirect to backend endpoint
                // The backend will handle the redirect to the final destination
                const backendUrl = `${import.meta.env.VITE_SERVER_HOST || 'http://localhost:3000'}/api/v1/url/${slug}`
                window.location.href = backendUrl
            } catch (err) {
                setError(err instanceof Error ? err.message : 'An error occurred')
                setLoading(false)
            }
        }

        handleRedirect()
    }, [slug])

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
                    <Box className="leaf-icon">🍂</Box>

                    <Typography variant="h4" component="h1" className="title">
                        Redirecting
                    </Typography>

                    <Typography variant="body1" className="subtitle">
                        {error ? 'An error occurred' : 'Please wait while we prepare your journey...'}
                    </Typography>

                    <Box className="progress-container">
                        {loading && !error && (
                            <CircularProgress
                                sx={{
                                    color: '#C87137',
                                    '& .MuiCircularProgress-circle': {
                                        strokeLinecap: 'round',
                                    },
                                }}
                            />
                        )}
                    </Box>

                    {error && (
                        <Typography variant="body2" className="error-message">
                            {error}
                        </Typography>
                    )}

                    <Typography variant="caption" className="description">
                        {loading ? 'Validating your request...' : 'Please try again or contact support'}
                    </Typography>
                </Box>
            </Container>
        </Box>
    )
}
