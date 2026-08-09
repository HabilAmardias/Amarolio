import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { HomePage } from '../pages/home';
import { Navbar } from '../components/Navbar';
import { Box } from '@mui/material';
import { LoginPage } from '../pages/login';
import { DashboardPage } from '../pages/dashboard';
import { URLStatsPage } from '../pages/dashboard/URLStatsPage';
import { ProtectedRoute } from '../components/ProtectedRoute';
import { HelmetProvider } from 'react-helmet-async';
import { QRISPage } from '../pages/donation/qris';
import { Error404Page } from '../pages/error';

export function AppRouter() {
  return (
    <HelmetProvider>
      <BrowserRouter>
        <Box sx={{ minHeight: '100vh', display: 'flex', flexDirection: 'column', position: 'relative' }}>
          <Navbar />
          <Box component="main" sx={{ flexGrow: 1 }}>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/login" element={<LoginPage />} />
              <Route
                path="/dashboard"
                element={
                  <ProtectedRoute>
                    <DashboardPage />
                  </ProtectedRoute>
                }
              />
              <Route path="/donation/qris" element={
                <QRISPage />
              }></Route>
              <Route
                path="/dashboard/url/:id"
                element={
                  <ProtectedRoute>
                    <URLStatsPage />
                  </ProtectedRoute>
                }
              />
              <Route path="*" element={<Error404Page />} />
            </Routes>
          </Box>
        </Box>
      </BrowserRouter>
    </HelmetProvider>
  );
}
