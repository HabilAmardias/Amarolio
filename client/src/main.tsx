import React from "react";
import ReactDOM from "react-dom/client";
import { Provider as JotaiProvider } from "jotai";
import ThemeProvider from "./theme/ThemeProvider";
import AppRouter from "./router";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <JotaiProvider>
      <ThemeProvider>
        <AppRouter />
      </ThemeProvider>
    </JotaiProvider>
  </React.StrictMode>
);
