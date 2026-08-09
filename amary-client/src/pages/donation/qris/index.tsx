import { Box, Paper, Typography } from '@mui/material'
import qrisImage from '../../../assets/qris.JPG'

export function QRISPage() {
    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                minHeight: 'calc(100vh - 64px)',
                padding: { xs: 2, sm: 4 },
                boxSizing: 'border-box',
            }}
        >
            <Typography
                variant="h4"
                component="h1"
                sx={{
                    color: 'primary.main',
                    fontWeight: 700,
                    mb: 3,
                    textAlign: 'center',
                }}
            >
                Support Amary
            </Typography>
            <Paper
                elevation={0}
                variant="outlined"
                sx={{
                    p: { xs: 2, sm: 3 },
                    borderRadius: 4,
                    display: 'flex',
                    justifyContent: 'center',
                }}
            >
                <Box
                    component="img"
                    src={qrisImage}
                    alt="QRIS Code"
                    sx={{
                        maxWidth: '100%',
                        maxHeight: '65vh',
                        width: 'auto',
                        height: 'auto',
                        objectFit: 'contain',
                        borderRadius: 2,
                    }}
                />
            </Paper>
        </Box>
    )
}