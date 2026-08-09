import { Navigate } from 'react-router-dom';
import { CircularProgress, Box } from '@mui/material';
import { useAuth } from '../controllers/useAuth';

interface ProtectedRouteProps {
  children: React.ReactNode;
}

export function ProtectedRoute({ children }: ProtectedRouteProps) {
  const { user, isLoading } = useAuth()

  if (isLoading) {
    return (
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          height: '100vh',
          background: 'radial-gradient(1100px 560px at 8% -8%, #ffe8c2 0%, transparent 55%), radial-gradient(900px 560px at 100% 4%, #f9c08a 0%, transparent 52%), radial-gradient(1000px 640px at 55% 112%, #f3a45b 0%, transparent 55%), linear-gradient(180deg, #fff8ef 0%, #f8e7d2 100%)',
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

  if (!user) {
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
}
