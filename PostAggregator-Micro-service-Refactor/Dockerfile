# Use the correct Go version that matches your go.mod
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build gRPC server
RUN CGO_ENABLED=0 GOOS=linux go build -o /grpc-server ./cmd/grpc-server

# Build GraphQL server
RUN CGO_ENABLED=0 GOOS=linux go build -o /graphql-server ./cmd/graphql-server

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install dependencies
RUN apk --no-cache add ca-certificates

# Copy binaries from builder
COPY --from=builder /grpc-server /app/grpc-server
COPY --from=builder /graphql-server /app/graphql-server

# Copy data file
COPY data.json /app/data.json

EXPOSE 50051 8080

# Default command (can be overridden in compose)
CMD ["/app/grpc-server"]