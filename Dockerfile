FROM golang:1.23.4-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go source code and .env file to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Create a new stage for the final image
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/main /
# Copy the .env file into the root directory of the final image
COPY --from=builder /app/.env /

# Set the command to run the binary
CMD ["/main"]