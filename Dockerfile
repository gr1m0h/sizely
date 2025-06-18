# Build stage
FROM golang:1.21-alpine AS builder

# Install git and ca-certificates (needed for downloading dependencies)
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o sizely \
    ./cmd/sizely

# Final stage
FROM scratch

# Copy ca-certificates for HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary
COPY --from=builder /app/sizely /sizely

# Copy example files
COPY --from=builder /app/examples /examples

# Set the binary as the entrypoint
ENTRYPOINT ["/sizely"]

# Default command
CMD ["--help"]

# Labels
LABEL org.opencontainers.image.title="sizely"
LABEL org.opencontainers.image.description=""
LABEL org.opencontainers.image.vendor="gr1m0h"
LABEL org.opencontainers.image.licenses="MIT"
