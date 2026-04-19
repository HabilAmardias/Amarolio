import React, { Suspense } from "react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { CircularProgress, Box } from "@mui/material";
import RootLayout from "../components/layout/RootLayout";

const HomeView = React.lazy(() => import("../views/Home/index"));
const ExperienceView = React.lazy(() => import("../views/Experience/index"));
const ProjectsView = React.lazy(() => import("../views/Projects/index"));

const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      { index: true, element: <HomeView /> },
      { path: "experience", element: <ExperienceView /> },
      { path: "projects", element: <ProjectsView /> },
    ],
  },
]);

export default function AppRouter() {
  return (
    <Suspense
      fallback={
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            minHeight: "100vh",
          }}
        >
          <CircularProgress />
        </Box>
      }
    >
      <RouterProvider router={router} />
    </Suspense>
  );
}