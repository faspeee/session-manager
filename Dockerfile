# Stage 1: Build the Go binary
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a minimal image with the Go binary
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has execute permissions
RUN chmod +x main

# Command to run the executable
CMD ["./main"]