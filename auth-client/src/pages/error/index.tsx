import { Box, Button, Paper, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import { AuthShell } from "../../components/auth/AuthShell";
import { Seo } from "../../components/common/Seo";

export function Error404Page() {
  return (
    <AuthShell subtitle="Not found">
      <Seo title="Not found | Amarolio Account" />
      <Paper variant="outlined" sx={{ p: { xs: 3, sm: 4 }, textAlign: "center" }}>
        <Typography variant="h3" sx={{ fontWeight: 700, mb: 1 }}>
          404
        </Typography>
        <Typography variant="body1" color="text.secondary" sx={{ mb: 3 }}>
          The page you are looking for does not exist.
        </Typography>
        <Box>
          <Button variant="contained" component={RouterLink} to="/">
            Go to your account
          </Button>
        </Box>
      </Paper>
    </AuthShell>
  );
}
