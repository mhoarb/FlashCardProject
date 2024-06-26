#Stage 1
# Use the official Golang image as the base image
FROM golang:1.22-alpine As builder

# Set environment variables for proxychains
ENV GOPROXY=https://goproxy.io,direct


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directorccleay to the Working Directory inside the container
COPY . .

RUN go get

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .


# Stage 2: Create a minimal image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Set the entry point
ENTRYPOINT ["/main"]