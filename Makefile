.PHONY: build run upload deploy
build:
	env GOOS=js GOARCH=wasm \
	go build -o webpage/game.wasm github.com/GoWorkshopConference/golang-game/cmd/game
run:
	go run cmd/game/main.go
upload:
	# make upload BUCKET_NAME=your-bucket-name
	@if [ -z "$(BUCKET_NAME)" ]; then \
		echo "Error: BUCKET_NAME environment variable is required"; \
		echo "Usage: make upload BUCKET_NAME=your-bucket-name"; \
		exit 1; \
	fi
	aws s3 sync webpage/ s3://$(BUCKET_NAME)/ --delete
deploy: build upload
