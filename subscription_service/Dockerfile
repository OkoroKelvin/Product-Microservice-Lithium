# Build stage
FROM golang:1.23.2-alpine AS builder

# Set working directory
WORKDIR /app

# Install required system packages
RUN apk add --no-cache git protoc protobuf-dev make

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 

RUN export PATH="$PATH:$(go env GOPATH)/bin"

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Generate protobuf files
RUN make generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Final stage
FROM alpine:3.19

# Set working directory
WORKDIR /app

# Install necessary runtime packages
RUN apk add --no-cache ca-certificates

RUN ls -la /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8081

# Run the application
CMD ["/app/main"]
