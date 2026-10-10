import { Box, Container } from "@mui/material";
import { Brand } from "./Brand";
import { ThemeToggle } from "./ThemeToggle";

interface AuthShellProps {
  children: React.ReactNode;
  subtitle?: string;
  width?: number;
}

export function AuthShell({
  children,
  subtitle = "Account",
  width = 460,
}: AuthShellProps) {
  return (
    <Box
      sx={{
        minHeight: "100vh",
        display: "flex",
        flexDirection: "column",
      }}
    >
      <Box
        component="header"
        sx={{
          px: { xs: 2, sm: 4 },
          py: 2,
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
        }}
      >
        <Brand subtitle={subtitle} />
        <ThemeToggle />
      </Box>

      <Container
        maxWidth="sm"
        sx={{
          flexGrow: 1,
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          py: { xs: 3, sm: 6 },
        }}
      >
        <Box
          className="animate-fade-up"
          sx={{ width: "100%", maxWidth: width }}
        >
          {children}
        </Box>
      </Container>
    </Box>
  );
}
