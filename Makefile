.PHONY: run build clean

run:
	go build -o forum && ./forum

clean:
	rm -f forum

test:
	go test ./...
