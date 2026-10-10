import { useEffect } from 'react';
import { Box, CircularProgress } from '@mui/material';
import { useLocation, useNavigate } from 'react-router-dom';
import { useAtom } from 'jotai';
import { userModel } from '../../models/user/model';
import { redirectToLogin } from '../../lib/authRedirect';

interface LocationState {
  from?: {
    pathname?: string;
    search?: string;
  };
}

export function LoginPage() {
  const [user] = useAtom(userModel.userAtom);
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      navigate('/', { replace: true });
      return;
    }

    const state = location.state as LocationState | null;
    const from = state?.from;
    const returnTo = from?.pathname
      ? `${window.location.origin}${from.pathname}${from.search ?? ''}`
      : `${window.location.origin}/`;

    redirectToLogin(returnTo);
  }, [user, location.state, navigate]);

  return (
    <Box
      sx={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
      }}
    >
      <CircularProgress
        sx={{
          color: 'primary.main',
        }}
      />
    </Box>
  );
}
