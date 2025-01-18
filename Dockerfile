# Build stage
FROM golang:1.23.5-alpine AS builder

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Initialize module and download dependencies
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o calculator-api

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/calculator-api .

# Expose port 8080 (assuming this is the default port, adjust if different)
EXPOSE 8080

# Run the binary
CMD ["./calculator-api"]
