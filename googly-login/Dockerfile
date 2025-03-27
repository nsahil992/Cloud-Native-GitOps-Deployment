# Build Stage
FROM golang:1.23.1 AS builder
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application for Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Runtime Stage
FROM debian:bookworm-slim
WORKDIR /app

# Copy only the compiled binary and necessary files
COPY --from=builder /app/main .
COPY static/ static/
COPY .env .env

# Ensure the binary is executable
RUN chmod +x main

# Expose port 8080
EXPOSE 8080

# Start the application
CMD ["./main"]
