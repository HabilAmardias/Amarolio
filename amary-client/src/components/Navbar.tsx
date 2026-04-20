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
            flexGrow: 1,
            color: '#c25e00',
            fontWeight: 700,
          }}
        >
          <Link to="/" style={{ textDecoration: 'none', color: 'inherit' }}>
            Amary
          </Link>
        </Typography>
        <Box>
          {user ? (
            <>
              <Typography 
                variant="body2" 
                sx={{ 
                  mr: 2, 
                  display: 'inline-block',
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
