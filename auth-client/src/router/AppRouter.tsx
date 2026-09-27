import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import { HelmetProvider } from "react-helmet-async";
import { AccountPage } from "../pages/account";
import { LoginPage } from "../pages/login";
import { RegisterPage } from "../pages/register";
import { VerifyPage } from "../pages/verify";
import { LogoutPage } from "../pages/logout";
import { Error404Page } from "../pages/error";

export function AppRouter() {
  return (
    <HelmetProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<AccountPage />} />
          <Route path="/account" element={<Navigate to="/" replace />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/verify" element={<VerifyPage />} />
          <Route path="/logout" element={<LogoutPage />} />
          <Route path="*" element={<Error404Page />} />
        </Routes>
      </BrowserRouter>
    </HelmetProvider>
  );
}
