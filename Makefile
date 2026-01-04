.PHONY: all test build run clean

BINARY_NAME=date-cli.exe
ARGS=

all: test build

test:
	go test -v ./...

build:
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go $(ARGS)

clean:
	rm -f $(BINARY_NAME)
