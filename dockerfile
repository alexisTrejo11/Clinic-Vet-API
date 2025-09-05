# Stage 1: Use a Go image to build the application
FROM golang:1.23.0-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal image with the application binary
FROM alpine:3.18

# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .  # Copy environment file if needed

# Create necessary directories for your application
RUN mkdir -p logs

# Expose the port the app runs on (use 8000 to match your compose file)
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
