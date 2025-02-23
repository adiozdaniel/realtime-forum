.PHONY: run build clean

run:
	cd backend && go run main.go

build:
	go build -o myprogram

clean:
	rm -f myprogram
