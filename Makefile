APP=app

build:
	go build -o bin/$(APP) ./cmd/$(APP)

run:
	go run ./cmd/$(APP) test

test:
	go test ./...

clean:
	rm -rf bin
