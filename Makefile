build:
	go build -o bin/gochat

run: build
	./bin/gochat

clean:
	rm bin/gochat
