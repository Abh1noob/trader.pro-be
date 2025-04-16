# Use Go Alpine image as base
FROM golang:1.20-alpine as builder

# Set up the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o /app/backend .

# Start a new stage from the Alpine image to keep the final image lightweight
FROM alpine:latest

# Install necessary dependencies for running Go applications
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/backend .

# Expose the port on which your app will run
EXPOSE 8080

# Run the Go application
CMD ["./backend"]
