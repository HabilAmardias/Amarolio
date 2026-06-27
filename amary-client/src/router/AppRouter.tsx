import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { HomePage } from '../pages/home';
import { Navbar } from '../components/Navbar';
import { Box } from '@mui/material';
import { LoginPage } from '../pages/login';
import { LoginCallbackPage } from '../pages/login/callback'
import { DashboardPage } from '../pages/dashboard';
import { URLStatsPage } from '../pages/dashboard/URLStatsPage';
import { ProtectedRoute } from '../components/ProtectedRoute';
import { HelmetProvider } from 'react-helmet-async';

export function AppRouter() {
  return (
    <HelmetProvider>
      <BrowserRouter>
        <Box sx={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
          <Navbar />
          <Box component="main" sx={{ flexGrow: 1 }}>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/login/callback" element={<LoginCallbackPage />} />
              <Route
                path="/dashboard"
                element={
                  <ProtectedRoute>
                    <DashboardPage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="/dashboard/url/:id"
                element={
                  <ProtectedRoute>
                    <URLStatsPage />
                  </ProtectedRoute>
                }
              />
            </Routes>
          </Box>
        </Box>
      </BrowserRouter>
    </HelmetProvider>
  );
}
