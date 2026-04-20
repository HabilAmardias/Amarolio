import { Navigate } from 'react-router-dom';
import { useAtom } from 'jotai';
import { authAtom, authLoadingAtom } from '../models/user.model';
import { CircularProgress, Box } from '@mui/material';

interface ProtectedRouteProps {
  children: React.ReactNode;
}

export function ProtectedRoute({ children }: ProtectedRouteProps) {
  const [user] = useAtom(authAtom);
  const [isLoading] = useAtom(authLoadingAtom);

  if (isLoading) {
    return (
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          height: '100vh',
          background: 'radial-gradient(ellipse at center, #1a1a2e 0%, #0a0a0f 100%)',
        }}
      >
        <CircularProgress 
          sx={{ 
            color: '#c25e00',
            boxShadow: '0 0 30px rgba(194, 94, 0, 0.6)',
          }} 
        />
      </Box>
    );
  }

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
}
