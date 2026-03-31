# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for go modules sometimes)
RUN apk add --no-cache git

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build binary
RUN go build -o vpn-manager cmd/server/main.go

# ---------- RUN STAGE ----------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/vpn-manager .

# Expose port
EXPOSE 8080

# Run binary
CMD ["./vpn-manager"]
