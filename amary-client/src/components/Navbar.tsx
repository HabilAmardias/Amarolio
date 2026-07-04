import { AppBar, Toolbar, Typography, Box, Button, IconButton } from '@mui/material';
import DashboardIcon from '@mui/icons-material/Dashboard';
import { Link } from 'react-router-dom';
import { useAuth } from '../controllers/useAuth';

export function Navbar() {
  const { user, logout } = useAuth();

  return (
    <AppBar position="static" color="default">
      <Toolbar
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
          flexWrap: 'nowrap',
          gap: 1,
        }}
      >
        {/* LEFT SECTION: Logo & Donate */}
        <Box
          sx={{
            display: 'flex',
            alignItems: 'center',
            flex: 1,
            justifyContent: 'flex-start',
            minWidth: 0,
          }}
        >
          <Typography
            variant="h6"
            component="div"
            sx={{
              color: '#c25e00',
              fontWeight: 700,
              whiteSpace: 'nowrap',
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
              ml: { xs: 1, sm: 2, md: 3 },
              color: '#ffffff',
              fontWeight: 600,
              fontSize: { xs: '0.75rem', sm: '0.9rem', md: '1rem' },
              whiteSpace: 'nowrap',
              minWidth: 'auto',
            }}
          >
            ☕ Donate
          </Button>
        </Box>

        {/* CENTER SECTION: Dashboard Button (Responsive) */}
        {user && (
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              flex: 1,
            }}
          >
            {/* Desktop View: Full Text Button */}
            <Button
              color="inherit"
              component={Link}
              to="/dashboard"
              sx={{
                display: { xs: 'none', sm: 'inline-flex' }, // Hidden on mobile
                color: '#ffffff',
                fontWeight: 600,
                fontSize: { sm: '0.9rem', md: '1rem' },
                whiteSpace: 'nowrap',
                '&:hover': {
                  backgroundColor: 'rgba(194, 94, 0, 0.08)',
                },
              }}
            >
              Dashboard
            </Button>

            {/* Mobile View: Clean Icon Button to prevent overlapping */}
            <IconButton
              component={Link}
              to="/dashboard"
              sx={{
                display: { xs: 'inline-flex', sm: 'none' }, // Only visible on mobile
                color: '#c25e00',
                '&:hover': {
                  backgroundColor: 'rgba(194, 94, 0, 0.08)',
                },
              }}
              aria-label="dashboard"
            >
              <DashboardIcon />
            </IconButton>
          </Box>
        )}

        {/* RIGHT SECTION: Auth Actions */}
        <Box
          sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'flex-end',
            gap: { xs: 1, sm: 2 },
            flex: 1,
          }}
        >
          {user ? (
            <>
              <Typography
                variant="body2"
                sx={{
                  color: '#5d4037',
                  display: { xs: 'none', md: 'block' },
                }}
              >
                {user.username}
              </Typography>
              <Button
                color="inherit"
                onClick={logout}
                sx={{
                  color: '#ffffff',
                  fontWeight: 600,
                  fontSize: { xs: '0.75rem', sm: '0.9rem', md: '1rem' },
                  px: { xs: 1, sm: 2 },
                  whiteSpace: 'nowrap',
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
                fontSize: { xs: '0.75rem', sm: '0.9rem', md: '1rem' },
                whiteSpace: 'nowrap',
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