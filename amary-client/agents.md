URL Shortener Web Client — Detailed Context Document
Project Overview
A responsive, minimalist URL shortener web client built with React, TypeScript, Jotai, Material UI, and React Router. The client supports anonymous URL shortening (with a fixed 24-hour expiration) and authenticated shortening (with user-defined or no expiration). All API calls are mocked and clearly marked for replacement.

Tech Stack
Layer                       Library 
UI Framework                React 18
Language                    TypeScript
State Management            Jotai
Component Library           Material UI (MUI)
Routing                     React Router
Architecture                MVC + folder-based routing

Architecture
MVC Mapping

Model — Jotai atoms + TypeScript interfaces (data shape, global state)
View — React components (pages + shared UI components)
Controller — Custom hooks (useShorten, useAuth) that orchestrate state mutations and API calls

Folder Structure
src/
├── api/                        # Mock API layer (replace internals with real calls)
│   ├── auth.api.ts             # login(), logout(), getMe()
│   └── url.api.ts              # shortenUrl(), getUserUrls()
│
├── models/                     # TypeScript interfaces and Jotai atoms
│   ├── user.model.ts           # User interface + authAtom + loadingAtom
│   └── url.model.ts            # ShortenRequest, ShortenResponse interfaces + urlHistoryAtom
│
├── controllers/                # Business logic hooks (MVC Controller layer)
│   ├── useAuth.ts              # login, logout, restoreSession handlers
│   └── useShorten.ts           # handleShorten, form state, validation logic
│
├── components/                 # Reusable, dumb UI components
│   ├── Navbar.tsx
│   ├── ShortenForm.tsx
│   ├── ResultCard.tsx
│   └── ProtectedRoute.tsx
│
├── pages/                      # Folder-based routing — each folder = one route
│   ├── home/
│   │   └── index.tsx           # Route: /
│   ├── login/
│   │   └── index.tsx           # Route: /login
│   └── dashboard/
│       └── index.tsx           # Route: /dashboard (protected)
│
├── router/
│   └── AppRouter.tsx           # Central route definitions
│
├── theme/
│   └── theme.ts                # MUI minimalist theme config
│
└── main.tsx                    # App entry point with Provider setup

Models (Data Layer)
user.model.ts
tsexport interface User {
  id: string;
  email: string;
  name: string;
}

// Jotai atoms
export const authAtom = atom<User | null>(null);
export const authLoadingAtom = atom<boolean>(true); // true while restoring session
url.model.ts
tsexport interface ShortenRequest {
  originalUrl: string;
  expiresInDays: number | null; // null = no expiration (auth only)
}

export interface ShortenResponse {
  shortUrl: string;
  originalUrl: string;
  expiresAt: string | null; // ISO date string or null
  createdAt: string;
}

// Jotai atom
export const urlHistoryAtom = atom<ShortenResponse[]>([]);

API Layer (Mock — Replace Internals)
auth.api.ts
ts// MOCK — replace function bodies with real HTTP calls

export async function login(email: string, password: string): Promise<User> {
  // MOCK: simulate successful login
  return { id: "1", email, name: "Demo User" };
}

export async function logout(): Promise<void> {
  // MOCK: clear session cookie/token
}

export async function getMe(): Promise<User | null> {
  // MOCK: return null (no active session)
  return null;
}
url.api.ts
ts// MOCK — replace function bodies with real HTTP calls

export async function shortenUrl(req: ShortenRequest): Promise<ShortenResponse> {
  // MOCK: simulate shortening
  return {
    shortUrl: `https://short.ly/${Math.random().toString(36).slice(2, 8)}`,
    originalUrl: req.originalUrl,
    expiresAt: req.expiresInDays
      ? new Date(Date.now() + req.expiresInDays * 86400000).toISOString()
      : null,
    createdAt: new Date().toISOString(),
  };
}

export async function getUserUrls(): Promise<ShortenResponse[]> {
  // MOCK: return empty history
  return [];
}

Controllers (Business Logic Hooks)
useAuth.ts
Responsibilities:

Reads/writes authAtom and authLoadingAtom
Calls login(), logout(), getMe() from the API layer
Exposes: { user, isLoading, login, logout }
On mount, calls getMe() to restore session (sets authLoadingAtom to false when done)

useShorten.ts
Responsibilities:

Manages local form state: url (string), expiresInDays (number | null), result, error, isLoading
Reads authAtom to determine if user is authenticated
If unauthenticated: forces expiresInDays = 1, hides expiry controls
If authenticated: exposes expiry input and a "No expiration" toggle
Validates URL format before calling shortenUrl()
On success, prepends result to urlHistoryAtom
Exposes: { url, setUrl, expiresInDays, setExpiresInDays, noExpiry, setNoExpiry, result, error, isLoading, handleShorten }


Pages & Routing
Route Table
PathComponentAuth Required/pages/home/index.tsxNo/loginpages/login/index.tsxNo (redirect if already logged in)/dashboardpages/dashboard/index.tsxYes
pages/home/index.tsx

Renders <ShortenForm /> centered on screen
If authenticated, shows a link to /dashboard
After shortening, renders <ResultCard /> below the form

pages/login/index.tsx

Email + password form
Calls useAuth().login()
On success, redirects to /
On failure, shows inline MUI Alert error

pages/dashboard/index.tsx

Protected by <ProtectedRoute />
Lists urlHistoryAtom entries in a MUI Table
Each row: short URL (clickable), original URL (truncated), expires at, created at
Has a <ShortenForm /> at the top so users can shorten without leaving

router/AppRouter.tsx
tsx<Routes>
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
</Routes>

Components
<Navbar />

MUI AppBar with minimalist styling (white background, subtle border bottom, no elevation)
Left: app name/logo
Right: if logged in → user email + Logout button; if not → Login button
Uses useAuth() for state and action

<ShortenForm />

MUI TextField for URL input (full-width)
If authenticated:

MUI TextField (type=number, min=1) for expiry days
MUI Checkbox or Switch labeled "No expiration" — when checked, disables and nulls the days field


MUI Button (full-width, contained) to submit
Inline CircularProgress on loading
Inline Alert on error

<ResultCard />

MUI Card with subtle border
Displays: short URL with copy-to-clipboard IconButton (uses navigator.clipboard), original URL, expiration info
On copy success: brief MUI Snackbar confirmation

<ProtectedRoute />

Reads authAtom and authLoadingAtom
While loading: renders centered <CircularProgress />
If no user: <Navigate to="/login" replace />
Otherwise: renders {children}


Theme (theme/theme.ts)
MUI theme configuration targeting a clean, minimalist aesthetic:
tscreateTheme({
  palette: {
    mode: "light",
    primary: { main: "#111111" },       // near-black primary
    secondary: { main: "#757575" },     // mid-grey secondary
    background: { default: "#FAFAFA", paper: "#FFFFFF" },
  },
  typography: {
    fontFamily: "'Inter', 'Helvetica Neue', Arial, sans-serif",
    button: { textTransform: "none", fontWeight: 600 },
  },
  shape: { borderRadius: 8 },
  components: {
    MuiButton: {
      defaultProps: { disableElevation: true },
    },
    MuiAppBar: {
      styleOverrides: {
        root: { boxShadow: "none", borderBottom: "1px solid #E0E0E0" },
      },
    },
    MuiTextField: {
      defaultProps: { size: "small", variant: "outlined" },
    },
  },
})

Lighthouse & Performance Targets
To hit ≥90 Lighthouse score:
ConcernImplementationNo layout shiftFixed-height form container, no dynamic height jumpsSemantic HTMLUse <main>, <nav>, <header>, proper heading hierarchyAccessible labelsAll TextField have label prop; buttons have descriptive textColor contrastNear-black on white palette passes WCAG AANo render-blockingMUI loaded via tree-shaking; Inter font via <link rel="preconnect"> in index.htmlResponsiveMUI Grid/Box with sx breakpoints; maxWidth="sm" container on home pageImage-freePure text + icon UI, no unoptimized images

State Flow Summary
User visits /
  → authLoadingAtom = true → ProtectedRoute shows spinner (dashboard only)
  → getMe() resolves → authAtom = null or User → authLoadingAtom = false

User types URL + clicks Shorten
  → useShorten.handleShorten()
  → if unauthenticated: expiresInDays forced to 1
  → shortenUrl(req) called → result set → urlHistoryAtom prepended

User logs in at /login
  → useAuth.login() → authAtom updated → navigate("/")

User visits /dashboard without auth
  → ProtectedRoute → <Navigate to="/login" />