.PHONY: build
build:
		go build -v ./cmd/application

.PHONY: deploybuild
deploybuild:
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/application
.DEFAULT_GOAL := build