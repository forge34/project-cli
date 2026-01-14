APP=app

build:
	go build -o bin/$(APP) ./cmd/$(APP)

run:
	go run ./cmd/$(APP) --template=express-ts create server

test:
	go test ./...

clean:
	rm -rf bin
