APP=pjc
BINARY=bin/$(APP)
TARGET=/usr/bin/$(APP)

build:
	go build -o $(BINARY) ./cmd/app

run:
	go run ./cmd/$(APP) --template=express-ts create server

test:
	go test ./...

clean:
	rm -rf bin

install: build
	@echo "Installing $(APP) to $(TARGET)..."
	sudo cp $(BINARY) $(TARGET)
	sudo chmod +x $(TARGET)
	@echo "$(APP) installed. You can now run it globally."
