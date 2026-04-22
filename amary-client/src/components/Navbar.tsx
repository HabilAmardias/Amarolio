import { AppBar, Toolbar, Typography, Button, Box } from '@mui/material';
import { useAuth } from '../controllers/useAuth';
import { Link } from 'react-router-dom';

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
        {user && (
          <Button 
            component={Link} 
            to="/dashboard"
            sx={{
              color: '#c25e00',
              fontWeight: 700,
              ml: 3,
              background: 'none',
              boxShadow: 'none',
              '&:hover': {
                backgroundColor: 'none',
                boxShadow: 'none',
              },
            }}
          >
            Dashboard
          </Button>
        )}
        <Box sx={{ flexGrow: 1 }} />
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
          {user ? (
            <>
              <Typography 
                variant="body2" 
                sx={{ 
                  color: '#5d4037',
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
