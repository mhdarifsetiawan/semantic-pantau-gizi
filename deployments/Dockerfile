# Stage 1: build
FROM golang:1.25-alpine AS builder

# Install git and ca-certificates (needed for go modules and HTTPS)
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary — CGO disabled for fully static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/bin/api ./cmd/api

# Stage 2: runtime
# Use distroless for minimal attack surface — no shell, no package manager
FROM gcr.io/distroless/static-debian12

# Copy timezone data and certs from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary only
COPY --from=builder /app/bin/api /api

# Run as non-root user (distroless default: nonroot uid 65532)
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/api"]