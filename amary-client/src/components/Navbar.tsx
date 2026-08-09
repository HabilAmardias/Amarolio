import { AppBar, Toolbar, Typography, Box, Button, IconButton, Popper, Grow, Paper, ClickAwayListener, MenuList, MenuItem } from '@mui/material';
import DashboardIcon from '@mui/icons-material/Dashboard';
import LocalCafeOutlinedIcon from '@mui/icons-material/LocalCafeOutlined';
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
    <AppBar position="static" color="inherit">
      <Toolbar
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
          flexWrap: 'nowrap',
          gap: 1,
          px: { xs: 1.5, sm: 2, md: 3 },
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
              color: 'primary.main',
              fontWeight: 700,
              whiteSpace: 'nowrap',
              fontFamily: "'Merriweather', 'Georgia', serif",
            }}
          >
            <Link to="/" style={{ textDecoration: 'none', color: 'inherit' }}>
              Amary
            </Link>
          </Typography>

          <Button
            ref={ref}
            onClick={() => setOpen((prev) => !prev)}
            startIcon={<LocalCafeOutlinedIcon />}
            sx={{
              background: "none",
              color: "primary.main",
              ml: { xs: 1, sm: 2, md: 3 },
              fontWeight: 600,
              fontSize: { xs: '0.8rem', sm: '0.9rem', md: '1rem' },
              minWidth: 'auto',
              border: "none",
              boxShadow: "none",
              "&:hover": {
                boxShadow: "none",
                border: "none",
                background: "rgba(180, 83, 9, 0.08)",
              },
            }}
          >
            Donate
          </Button>
          <Popper
            open={open}
            anchorEl={ref.current}
            placement='bottom-start'
            transition
            disablePortal
            sx={{ zIndex: 1300 }}
          >
            {({ TransitionProps }) => (
              <Grow
                {...TransitionProps}
                style={{
                  transformOrigin: "left-top"
                }}
              >
                <Paper
                  elevation={2}
                  sx={{ mt: 0.5, overflow: 'hidden' }}
                >
                  <ClickAwayListener onClickAway={handleClose}>
                    <MenuList
                      autoFocusItem={open}
                      id="composition-menu"
                      aria-labelledby="composition-button"
                      sx={{ py: 0.5 }}
                    >
                      <MenuItem sx={{ px: 1 }}>
                        <Button
                          href='/donation/qris'
                          sx={{
                            width: '100%',
                            justifyContent: 'flex-start',
                            color: 'text.primary',
                            background: "none",
                            border: "none",
                            borderRadius: 10,
                            "&:hover": {
                              background: "rgba(180, 83, 9, 0.08)",
                            }
                          }}
                        >
                          QRIS
                        </Button>
                      </MenuItem>
                      <MenuItem sx={{ px: 1 }}>
                        <Button
                          href={import.meta.env.VITE_DONATION_URI}
                          target='_blank'
                          rel='noopener noreferrer'
                          sx={{
                            width: '100%',
                            justifyContent: 'flex-start',
                            color: 'text.primary',
                            background: "none",
                            border: "none",
                            borderRadius: 10,
                            "&:hover": {
                              background: "rgba(180, 83, 9, 0.08)",
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
                display: { xs: 'none', sm: 'inline-flex' },
                color: 'text.secondary',
                fontWeight: 600,
                fontSize: { sm: '0.9rem', md: '1rem' },
                whiteSpace: 'nowrap',
                '&:hover': {
                  backgroundColor: 'rgba(180, 83, 9, 0.08)',
                  color: 'primary.main',
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
                display: { xs: 'inline-flex', sm: 'none' },
                color: 'primary.main',
                '&:hover': {
                  backgroundColor: 'rgba(180, 83, 9, 0.08)',
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
                  color: 'text.secondary',
                  display: { xs: 'none', md: 'block' },
                }}
              >
                {user.username}
              </Typography>
              <Button
                color="inherit"
                onClick={logout}
                sx={{
                  color: 'text.secondary',
                  fontWeight: 600,
                  fontSize: { xs: '0.75rem', sm: '0.9rem', md: '1rem' },
                  px: { xs: 1, sm: 2 },
                  whiteSpace: 'nowrap',
                  '&:hover': {
                    backgroundColor: 'rgba(180, 83, 9, 0.08)',
                    color: 'primary.main',
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
                color: 'text.secondary',
                fontWeight: 600,
                fontSize: { xs: '0.75rem', sm: '0.9rem', md: '1rem' },
                whiteSpace: 'nowrap',
                '&:hover': {
                  backgroundColor: 'rgba(180, 83, 9, 0.08)',
                  color: 'primary.main',
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
