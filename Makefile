.PHONY: build
build:
	go build -trimpath -ldflags="-s -w" -o sf6c.exe

.PHONY: dev
dev:
	go run main.go

.PHONY: combo
combo:
	go run cmd/main.go
