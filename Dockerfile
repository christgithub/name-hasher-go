# Start from the official Go image
FROM golang:1.24.1-alpine

# Set environment variables
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    APP_HOME=/app

# Install git (needed to download some Go modules)
RUN apk add --no-cache git

# Create app directory
WORKDIR $APP_HOME

# Copy Go modules first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary
RUN go build -o sftp_indexer main.go

# Run the binary
ENTRYPOINT ["./sftp_indexer"]