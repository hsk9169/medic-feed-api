APP_NAME = storage

build:
	go build -o $(APP_NAME) cmd/$(APP_NAME)/main.go
