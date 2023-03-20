# Build variables
APP_NAME=qpt-to-jpg-service
DOCKER_IMAGE=ashwinshirva/$(APP_NAME)
DOCKER_TAG=0.0.13

# Build the Go binary
.PHONY: build
build:
	# Here, we are setting the GOOS and GOARCH environment variables to linux and amd64 respectively to ensure that the binary runs on the Alpine Linux distribution with the x86-64 architecture.
	# We are also setting CGO_ENABLED=0 to ensure that the binary is built without any dependencies on C libraries.
	# The -ldflags flag is used to set linker flags to remove debugging information, reduce the binary size and statically link any dependencies.
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s -extldflags=-static" -o qpt-to-jpg-service .

# Build the Docker image
.PHONY: docker
docker:
	docker build --no-cache -t $(DOCKER_IMAGE):$(DOCKER_TAG) --progress=plain .

# Push the Docker image to a registry
.PHONY: push
push:
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

# Run the Docker container
.PHONY: docker-run
docker-run:
	# docker run --rm -it -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)
	docker run --privileged --rm -it -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG) /bin/sh
	#docker run --rm -it -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)


# Clean up the project
.PHONY: clean
clean:
	rm -f $(APP_NAME)

# Build everything
.PHONY: all
all: build docker push
