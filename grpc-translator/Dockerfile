# Use the official Golang image as the base image
FROM golang:1.19.7-alpine3.17 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the application
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN go build -o app .

# Create a new image based on a minimal Alpine image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the build image to the new image
COPY --from=build /app/app .

# Expose port
EXPOSE 8080

# Set the container command to run the application
CMD ["/app/app"]
