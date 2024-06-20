# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Install necessary tools including proxychains
RUN apk --no-cache add proxychains-ng

# Set up proxychains configuration
RUN echo 'socks5 127.0.0.1 10808' >> /etc/proxychains.conf

# Set environment variables for proxychains
ENV GOPROXY=https://goproxy.io,direct


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
