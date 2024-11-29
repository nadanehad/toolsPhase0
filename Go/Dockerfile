FROM golang:1.20-alpine

# Install bash (required by wait-for-it.sh)
RUN apk add --no-cache bash

WORKDIR /app

# Copy wait-for-it script and make it executable
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .


# Build the Go application
RUN go build -o main .

# Expose the backend port
EXPOSE 8080

# Run the backend server
CMD ["./main"]
