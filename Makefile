.PHONY: run build clean

run:
	go build -o forum && ./forum

clean:
	rm -f forum

test:
	go test ./...
docker:
	docker build -t forum . && docker run -d -p 4000:4000 --name forumcontainer forum
