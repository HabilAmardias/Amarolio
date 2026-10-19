# Amarolio

Monorepo for Muhammad Habil Amardias's portfolio and the services behind it.

## Apps

### Portfolio (`client/`)
React, TypeScript, Vite, Material UI, and Jotai. Pages for home, experience, and projects. Runs at amarolio.id.

### Amary URL shortener
Three parts:

- `amary/`: Go service on Gin. Stores long URLs, converts IDs into Base62 short codes, and records visits. PostgreSQL holds links and visit records; Redis caches lookups. Custom codes and expiry are supported.
- `amary-client/`: React client. Anonymous links expire after 24 hours; signed-in users choose their own expiry.
- `amary-redirect/`: React page that resolves a short slug and forwards the visitor.

### Auth service (`auth/`)
Go service on Fiber. Handles registration, email verification, password and OTP sign-in, token refresh, and password reset. The service hashes passwords with Argon2id. Access tokens last 15 minutes and refresh tokens last 7 days. PostgreSQL stores users and OTP challenges. Redis caches users and counts failed sign-in attempts; three failures lock the account for an hour.

### Auth client (`auth-client/`)
React app that centralizes sign-in. A service sends the user here with a `redirect_uri`, and the client returns them after password, OTP, or registration. Signed-in users see their account, a logout button, and links to the Amarolio apps. Runs at auth.amarolio.id.

### API gateway (`gateway/`)
Go service on Fiber v3. Fronts the clients and proxies to `amary` and `auth`. Verifies JWT cookies, checks Cloudflare Turnstile, and rate limits each caller: 5 requests per second per IP for anonymous traffic, 20 per second per account for signed-in users. The gateway reads the client IP from `X-Real-IP`.

## Tech stack

| Path | Language | Framework | Data |
|------|----------|-----------|------|
| client | TypeScript | React, Vite, Material UI | - |
| amary | Go | Gin | PostgreSQL, Redis |
| amary-client | TypeScript | React, Vite, Material UI | - |
| amary-redirect | TypeScript | React, Vite | - |
| auth-client | TypeScript | React, Vite, Material UI | - |
| auth | Go | Fiber v3 | PostgreSQL, Redis |
| gateway | Go | Fiber v3 | - |

## Hosts

| Host | App | Dev port |
|------|-----|----------|
| amarolio.id | portfolio | 5173 |
| amary.amarolio.id | amary-client | 4173 |
| amary.id | amary-redirect | 8000 |
| auth.amarolio.id | auth-client | 4000 |
| api.amarolio.id | gateway | 3000 |

## Auth flow

1. A client redirects the user to `auth.amarolio.id` with a `redirect_uri`.
2. Registration goes through `/register`, and the service emails a verification link. `/verify` activates the account.
3. Sign-in starts at `/prelogin` with the email and password. The service emails a one-time code, and `/login` exchanges it for access and refresh cookies.
4. Password reset starts at `/reset-password/send`, which emails a link. `/reset-password` stores the new password and clears the session cookies.
5. The auth client checks `redirect_uri` against a configured allowlist, then sends the user back.

## Development

Docker Compose runs the stack. The Makefile wraps the common commands.

```bash
make build-amarolio   # gateway and portfolio
make build-amary      # shortener service, client, redirect
make build-auth       # auth service, stores, auth client
```

`docker-compose.dev.yaml` builds images from local source. `docker-compose.yaml` pulls prebuilt images. Use `make start-*` and `make stop-*` to run each app.
