# Stage 1: Builder (Debian-based for CGO support)
FROM golang:1.23.4 AS builder

# Enable CGO
ENV CGO_ENABLED=1

# Install required build tools for go-sqlite3
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build main app
RUN go build -o main .

# Build migration binary
RUN go build -o migrate internal/migration/app_db_migration.go

# Stage 2: Minimal runtime image (use `distroless` or slim Alpine with CGO support)
FROM debian:bookworm-slim

# Create non-root user (optional security)
RUN useradd -m appuser

# Copy binaries and env file
COPY --from=builder /app/main /
COPY --from=builder /app/migrate /migrate
COPY --from=builder /app/.env /

# Optional: Copy SSL certs if needed
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Switch to non-root
USER appuser

CMD ["/main"]
