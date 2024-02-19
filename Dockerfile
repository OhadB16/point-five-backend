# Use the official Golang image to create a build artifact.
FROM golang:1.18-alpine AS builder

# Install gcc, musl-dev, and sqlite-dev (for go-sqlite3)
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app and place the binary in /app
RUN go build -o /app/github-events

# Use the official Alpine image for a lean production container.
FROM alpine:latest  

WORKDIR /root/

# Copy the compiled binary from the builder stage.
COPY --from=builder /app/github-events .

# Install SQLite
RUN apk add --no-cache sqlite

# Run the binary
CMD ["/root/github-events"]
