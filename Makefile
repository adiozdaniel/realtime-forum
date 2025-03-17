.PHONY: run build clean test docker docker-clean

# Build and run the application
run:
	go build -o forum && ./forum

# Clean up the binary
clean:
	rm -f forum

# Run tests
test:
	go test ./...

# Build and run the Docker container
docker:
	docker build -t forum .
	docker run -d -p 4000:4000 --name forumcontainer forum

# Stop and remove the Docker container
docker-clean:
	docker stop forumcontainer || true
	docker rm forumcontainer || true
