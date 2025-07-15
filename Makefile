.PHONY: build run clean test

build:
	go build -o rm 

run:
	go run .

clean:
	rm -f rm
