# Stage 1: Use a Go image to build the application
FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install postgresql-client for health checks
RUN apk add --no-cache postgresql-client

# Copy the rest of your application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal image with the application binary
FROM alpine:3.18

# Install CA certificates for HTTPS connections and postgresql-client
RUN apk --no-cache add ca-certificates postgresql-client

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/db/migrations ./db/migrations/
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Create necessary directories for your application
RUN mkdir -p logs

# Copy and make executable the entrypoint script
COPY scripts/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose the port the app runs on
EXPOSE 8000

# Use entrypoint script instead of direct CMD
ENTRYPOINT ["/entrypoint.sh"]
CMD ["./main"]