FROM golang:1.23-alpine

WORKDIR /app

# Copy and install Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the whole app
COPY . .

# Build
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
