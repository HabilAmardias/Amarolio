import { AppBar, Toolbar, Typography, Box, Button, IconButton, Popper, Grow, Paper, ClickAwayListener, MenuList, MenuItem } from '@mui/material';
import DashboardIcon from '@mui/icons-material/Dashboard';
import { Link } from 'react-router-dom';
import { useAuth } from '../controllers/useAuth';
import { useRef, useState } from 'react';

export function Navbar() {
  const { user, logout } = useAuth();
  const [open, setOpen] = useState<boolean>(false)
  const ref = useRef<HTMLButtonElement>(null)

  const handleClose = (event: Event | React.SyntheticEvent) => {
    if (
      ref.current &&
      ref.current.contains(event.target as HTMLElement)
    ) {
      return;
    }

    setOpen(false);
  };

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
            ref={ref}
            onClick={() => setOpen((prev) => !prev)}
            sx={{
              background: "none",
              color: "#c25e00",
              ml: { xs: 1, sm: 2, md: 3 },
              fontWeight: 900,
              fontSize: { xs: '0.75rem', sm: '0.9rem', md: '1rem' },
              minWidth: 'auto',
              border: "none",
              boxShadow: "none",
              "&:hover": {
                boxShadow: "none",
                border: "none"
              }
            }}
          >
            Donate ☕
          </Button>
          <Popper
            open={open}
            anchorEl={ref.current}
            placement='bottom-start'
            transition
            disablePortal
          >
            {({ TransitionProps }) => (
              <Grow
                {...TransitionProps}
                style={{
                  transformOrigin: "left-top"
                }}
              >
                <Paper>
                  <ClickAwayListener onClickAway={handleClose}>
                    <MenuList
                      autoFocusItem={open}
                      id="composition-menu"
                      aria-labelledby="composition-button"
                      sx={{
                        bgcolor: "#c25e00"
                      }}
                    >
                      <MenuItem>
                        <Button
                          href='/donation/qris'
                          sx={{
                            background: "none",
                            border: "none",
                            "&:hover": {
                              background: "none",
                              border: "none"
                            }
                          }}>QRIS</Button>
                      </MenuItem>
                      <MenuItem>
                        <Button
                          href={import.meta.env.VITE_DONATION_URI}
                          target='_blank'
                          rel='noopener noreferrer'
                          sx={{
                            background: "none",
                            border: "none",
                            "&:hover": {
                              background: "none",
                              border: "none"
                            }
                          }}
                        >
                          Ko-Fi
                        </Button>
                      </MenuItem>
                    </MenuList>
                  </ClickAwayListener>
                </Paper>
              </Grow>
            )}
          </Popper>
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