# Build stage
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o guardian-grid-agent ./cmd/agent/main.go

# Final stage
FROM alpine:latest

# Install basic tools for scanning if needed
RUN apk add --no-cache ca-certificates procps iproute2

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/guardian-grid-agent .

# Ensure the agent has execution permissions
RUN chmod +x guardian-grid-agent

# Set the default API URL (can be overridden at runtime)
ENV API_URL=http://127.0.0.1:8080

# Run the agent
CMD ["./guardian-grid-agent"]
