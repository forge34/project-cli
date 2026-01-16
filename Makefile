APP=pjc
BINARY=bin/$(APP)
TARGET=/usr/bin/$(APP)

build:
	go build -o $(BINARY) ./cmd/app

run:
	go run ./cmd/app create express-ts server

test:
	go test ./...

clean:
	rm -rf bin


silent-build:
	@go build -o $(BINARY) ./cmd/app

install:
	@echo "Installing $(APP) to $(TARGET)..."
	@$(MAKE) silent-build
	@sudo cp $(BINARY) $(TARGET)
	@sudo chmod +x $(TARGET)
	@echo "$(APP) installed. You can now run it globally."

