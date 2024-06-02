# Build stage
FROM golang:1.22.3-bullseye AS build-stage
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire source directory
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/main.go

# Final stage
FROM debian:bullseye
WORKDIR /app

# Install PostgreSQL client
RUN apt-get update && apt-get install -y postgresql-client ca-certificates

# Copy the Go runtime from the build-stage
COPY --from=build-stage /usr/local/go/ /usr/local/go/

# Set the environment variable for Go
ENV PATH="/usr/local/go/bin:${PATH}"

# Copy the application binary and the source code from the build-stage
COPY --from=build-stage /api /api
COPY --from=build-stage /app /app

EXPOSE 8080

# Set the entrypoint to the application
ENTRYPOINT ["/bin/sh", "-c", "/api"]
