import { AppBar, Toolbar, Typography, Box, Button } from '@mui/material';
import { Link } from 'react-router-dom';
import { useAuth } from '../controllers/useAuth';

export function Navbar() {
  const { user, logout } = useAuth();

  return (
    <AppBar position="static" color="default">
      <Toolbar>
        <Typography
          variant="h6"
          component="div"
          sx={{
            color: '#c25e00',
            fontWeight: 700,
          }}
        >
          <Link to="/" style={{ textDecoration: 'none', color: 'inherit' }}>
            Amary
          </Link>
        </Typography>
        <Button
          href={import.meta.env.VITE_DONATION_URI}
          target="_blank"
          rel="noopener noreferrer"
          sx={{
            ml: 3,
            color: '#ffffff',
            fontWeight: 600,
            '&:hover': {
              backgroundColor: 'rgba(194, 94, 0, 0.08)',
            },
          }}
        >
          ☕ Donate
        </Button>
        <Box sx={{ flexGrow: 1 }} />
        <Box sx={{ display: 'flex', alignItems: 'center', gap: { xs: 1, sm: 2 } }}>
          {user ? (
            <>
              <Typography
                variant="body2"
                sx={{
                  color: '#5d4037',
                  display: { xs: 'none', sm: 'block' },
                }}
              >
                {user.email}
              </Typography>
              <Button
                color="inherit"
                onClick={logout}
                sx={{
                  color: '#ffffff',
                  fontWeight: 600,
                  fontSize: { xs: '0.75rem', sm: '1rem' },
                  px: { xs: 1, sm: 2 },
                  '&:hover': {
                    backgroundColor: 'rgba(194, 94, 0, 0.08)',
                  },
                }}
              >
                Logout
              </Button>
            </>
          ) : (
            <Button
              color="inherit"
              component={Link}
              to="/login"
              sx={{
                color: '#ffffff',
                fontWeight: 600,
                '&:hover': {
                  backgroundColor: 'rgba(194, 94, 0, 0.08)',
                },
              }}
            >
              Login
            </Button>
          )}
        </Box>
      </Toolbar>
    </AppBar>
  );
}
