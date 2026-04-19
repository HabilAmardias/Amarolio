import { AppBar, Toolbar, Typography, Button, Box } from "@mui/material";
import { NavLink } from "react-router-dom";

export default function Navbar() {
  const navLinks = [
    { to: "/", label: "Home" },
    { to: "/experience", label: "Experience" },
    { to: "/projects", label: "Projects" },
  ];

  return (
    <AppBar position="sticky" color="transparent" elevation={0}>
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1, fontWeight: 700 }}>
          Amarolio
        </Typography>
        <nav aria-label="Main navigation">
          <Box sx={{ display: "flex", gap: 1 }}>
            {navLinks.map((link) => (
              <Button
                key={link.to}
                component={NavLink}
                to={link.to}
                color="inherit"
                aria-current="page"
                sx={() => ({
                  color: "primary.main" ,
                })}
              >
                {link.label}
              </Button>
            ))}
          </Box>
        </nav>
      </Toolbar>
    </AppBar>
  );
}