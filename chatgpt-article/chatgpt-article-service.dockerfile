# Build stage
FROM golang:1.23 AS builder

WORKDIR /app
COPY . .

# Build with CGO_ENABLED=1
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./chatgptArticleService ./cmd/consumer

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary libraries for CGO
RUN apk add --no-cache libc6-compat gcc g++ make

# Copy the built binary from the builder stage
COPY --from=builder /app/chatgptArticleService /app/chatgptArticleService

COPY --from=builder /app/internal/schemas /app/schemas

# Ensure the binary has execute permissions
RUN chmod +x /app/chatgptArticleService

CMD ["/app/chatgptArticleService"]
