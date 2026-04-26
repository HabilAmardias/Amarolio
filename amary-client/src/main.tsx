import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'jotai';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme } from './theme/theme';
import { AppRouter } from './router/AppRouter';
import { GoogleOAuthProvider } from '@react-oauth/google';
 

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>    
    <GoogleOAuthProvider clientId={import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_ID}>
      <Provider>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <AppRouter />
        </ThemeProvider>
      </Provider>
    </GoogleOAuthProvider>
  </React.StrictMode>
);
