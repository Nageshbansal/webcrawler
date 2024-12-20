# Dockerfile for Go web crawler project

# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o webcrawler main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory in the container
WORKDIR /root/

# Copy the binary from the builder stage to the working directory in the container
COPY --from=builder /app/webcrawler .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./webcrawler"]
