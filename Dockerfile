FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o calculator-api

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/calculator-api .

EXPOSE 8080

CMD ["./calculator-api"]
