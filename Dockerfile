# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine AS build

# Set the working directory to /go-webserver
WORKDIR /go-webserver

# Copy the server code and key-value file to the container
COPY . .


# Build the Go application inside the container
RUN go mod tidy && go build -o server

# Use a smaller, Alpine-based image as the base image for the final container
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the server binary from the build container to the final container
COPY --from=0 /go-webserver/server .

EXPOSE 3000

# Set the default command to run the server binary
CMD ["./server", "-message", "Welcome to the Go web server!"]
