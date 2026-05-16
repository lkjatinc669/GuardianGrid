# --- Build Stage ---
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build production binary
# -s -w: Strips debug info and symbol tables (reduces size by ~30%)
# CGO_ENABLED=0: Ensures a fully static binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o /guardian-grid-agent \
    ./cmd/agent/main.go

# --- Final Stage ---
FROM alpine:3.19

# Install runtime utilities for Linux scanning
# procps: for process info (/proc)
# iproute2: for detailed network stats
RUN apk add --no-cache \
    ca-certificates \
    procps \
    iproute2 \
    tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /guardian-grid-agent .

# Metadata and permissions
RUN chmod +x guardian-grid-agent

# Environment defaults
ENV API_URL=http://127.0.0.1:8080
ENV GIN_MODE=release

# Note: The agent typically requires --privileged or root access 
# to scan the host's /proc and /sys directories.
ENTRYPOINT ["./guardian-grid-agent"]
