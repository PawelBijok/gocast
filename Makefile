PROJECT_NAME := "gocast"
PKG := "github.com/pafello/${PROJECT_NAME}"

.PHONY: build build-mac build-linux build-windows


build: 
	@go build -o bin/${PROJECT_NAME} -v cmd/*.go
build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o bin/${PROJECT_NAME}_mac -v cmd/*.go
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/${PROJECT_NAME}_linux -v cmd/*.go
build-windows:
	@GOOS=windows GOARCH=amd64 go build -o bin/${PROJECT_NAME}_windows.exe -v cmd/*.go
