if exist "bin" (
	cd bin
	rmdir /s /q .
	cd ..
)

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0


go build -ldflags="-s -w" -o bin/msp-core ./cmd/api/main.go
swag init -g cmd/api/main.go --parseDependency --output docs/


 