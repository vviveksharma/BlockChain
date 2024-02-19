# Dockerfile

# Stage 1: Build the Golang binary
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Stage 2: Create a minimal image with only the binary and necessary files
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/Makefile ./Makefile
# Install any additional dependencies here if needed


# Set environment variables
ENV DATABASE_URL "postgres://postgres:password@db:5432/mydatabase?sslmode=disable"
# Expose port 8000 for the Golang application
EXPOSE 8000

VOLUME /app/data

# Command to run the Golang application
CMD ["./app"]
