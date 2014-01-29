all: clean build test

build:
	go build -o terimg main.go

clean:
	rm -f terimg

test: build
	./terimg img/shipwreck.png 100

