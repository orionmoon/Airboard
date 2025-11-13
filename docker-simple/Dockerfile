# 🐳 Dockerfile Simple - Airboard All-in-One
# Une seule image contenant Frontend + Backend + Nginx
# Usage: docker build -t airboard .

# ===== STAGE 1: BUILD FRONTEND =====
FROM node:18-alpine AS frontend-builder

WORKDIR /app

# Copy package files
COPY frontend/package*.json ./
RUN npm ci --silent

# Copy source and build
COPY frontend/ .
ENV NODE_ENV=production
RUN npm run build

# ===== STAGE 2: BUILD BACKEND =====
FROM golang:1.21-alpine AS backend-builder

# Install dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go files
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy source and build
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o airboard-api .

# ===== STAGE 3: FINAL IMAGE =====
FROM nginx:alpine

# Install runtime dependencies
RUN apk add --no-cache ca-certificates curl supervisor

# Copy built files
COPY --from=frontend-builder /app/dist /usr/share/nginx/html
COPY --from=backend-builder /app/airboard-api /app/
COPY --from=backend-builder /app/init.sql /app/

# Copy configuration files
COPY docker-simple/nginx.conf /etc/nginx/conf.d/default.conf
COPY docker-simple/supervisor.conf /etc/supervisor/conf.d/airboard.conf

# Make backend executable
RUN chmod +x /app/airboard-api

# Environment variables (defaults - override in docker-compose)
ENV GIN_MODE=release
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_NAME=airboard
ENV DB_USER=airboard
# DB_PASSWORD should be set via docker-compose or runtime

# Health check
HEALTHCHECK --interval=30s --timeout=3s \
    CMD curl -f http://localhost/health || exit 1

# Expose port
EXPOSE 80

# Start services
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/airboard.conf"]