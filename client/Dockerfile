# Stage 1: Build
FROM oven/bun:latest AS builder

WORKDIR /app

COPY package*.json ./
RUN bun install

COPY . .
RUN bun build

# Stage 2: Serve with Nginx
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]