GOCMD=GO111MODULE=on go
BINARY_NAME=api

.PHONY: help
clean:
	rm -rf ./bin

build:
	
	swag fmt
	swag init -g cmd/api/main.go --parseDependency --output docs/ 
	@$(GOCMD) build -ldflags="-s -w" -o bin/${BINARY_NAME} cmd/api/main.go


	
