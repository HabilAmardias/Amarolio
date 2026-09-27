# Amarolio

Monorepo for the portfolio site of Muhammad Habil Amardias and the services that run behind it.

## Apps

### Portfolio site (`client/`)
React and TypeScript with Vite, Material UI, and Jotai. Three pages: home, experience, projects. Served from amarolio.id.

### Amary URL shortener
Amarolio ships a URL shortener with three parts:

- `amary/` — Go service built on Gin. It encrypts long URLs, turns IDs into short codes, and logs visits. PostgreSQL stores links and visit records; Redis caches lookups. Supports custom codes and link expiry.
- `amary-client/` — React client. Anonymous links expire after 24 hours; signed-in users set their own expiry.
- `amary-redirect/` — React page that resolves a short slug and sends the visitor onward.

### Auth service (`auth/`)
Go service built on Fiber. Users register and sign in with an email, password, and a one-time passcode emailed to them. The service issues JWT access and refresh tokens. PostgreSQL stores users; Redis caches sessions.

### Auth client (`auth-client/`)
React client that centralizes sign-in across Amarolio. Services redirect to it with a `redirect_uri`; after prelogin, OTP verification, or registration the client sends the user back. Signed-in users land on an account page that shows their profile, a logout button, and the list of Amarolio apps. Served from auth.amarolio.id.

### API gateway (`gateway/`)
Go service built on Fiber v3. It fronts the client apps and proxies to the amary and auth services. It enforces JWT auth, Cloudflare Turnstile checks, and rate limits.

## Tech stack

| Path | Language | Framework | Data |
|------|----------|-----------|------|
| client | TypeScript | React, Vite, Material UI | - |
| amary | Go | Gin | PostgreSQL, Redis |
| amary-client | TypeScript | React, Vite, Material UI | - |
| amary-redirect | TypeScript | React, Vite | - |
| auth-client | TypeScript | React, Vite, Material UI | - |
| auth | Go | Fiber v3 | PostgreSQL, Redis |
| gateway | Go | Fiber v3 | Redis |

## Development

Docker Compose runs the stack. The Makefile wraps the common commands.

```bash
make build-amarolio     # gateway + portfolio client
make build-amary        # shortener backend, client, redirect
make build-auth         # auth service and its stores
make build-auth-client  # centralized auth client
```

`docker-compose.dev.yaml` builds images from local source. `docker-compose.yaml` pulls prebuilt images from the registry. Use `make start-*` and `make stop-*` to manage each app.

## Deployment

GitHub Actions builds every service on pushes to `master` and pushes the images to GitHub Container Registry under `ghcr.io/habilamardias`. The production compose file pulls those images.
