.PHONY: build run
build:
	env GOOS=js GOARCH=wasm \
	go build -o webpage/game.wasm github.com/GoWorkshopConference/golang-game/cmd/game
run:
	go run cmd/game/main.go
