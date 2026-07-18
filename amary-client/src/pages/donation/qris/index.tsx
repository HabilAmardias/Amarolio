import { Box } from '@mui/material'
import qrisImage from '../../../assets/qris.JPG'

export function QRISPage() {
    return (
        <Box
            sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                minHeight: 'calc(100vh - 64px)',
                padding: { xs: 2, sm: 4 },
                boxSizing: 'border-box',
            }}
        >
            <Box
                component="img"
                src={qrisImage}
                alt="QRIS Code"
                sx={{
                    maxWidth: '100%',
                    maxHeight: '75vh',
                    width: 'auto',
                    height: 'auto',
                    objectFit: 'contain',
                    borderRadius: 2,
                    boxShadow: 3,
                }}
            />
        </Box>
    )
}