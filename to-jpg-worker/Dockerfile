# The purpose of the first stage is simply to copy the 
# application source code and the go.mod and go.sum files into
# the container. This is still useful because it isolates the
# application source code and Go module dependencies from the
# host machine and provides a consistent environment for building the application.

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

# The second stage then copies the compiled binary from the
# first stage and creates a new image that only contains the
# binary and minimal dependencies required to run the application.
# This reduces the size of the final Docker image and makes it
# easier to deploy the application to production.

# Create a new image based on a minimal Alpine image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the build image to the new image
COPY --from=build /app/qpt-to-jpg-service .

# Expose port
EXPOSE 8080

# Set the container command to run the application
CMD ["/app/qpt-to-jpg-service"]

