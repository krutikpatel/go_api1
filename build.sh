#!/bin/bash

# Build the application
build() {
    go build -o bin/app main.go
}

# Run the application
run() {
    go run main.go
}

# Clean build files
clean() {
    rm -rf bin/
}

# Run tests
test() {
    go test -v ./...
}

# Run linter
lint() {
    golangci-lint run
}

# Generate swagger docs (if using swagger)
swagger() {
    swag init -g main.go
}

# Build for multiple platforms
build_all() {
    GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64 main.go
    ## GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64.exe cmd/server/main.go
    ## GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64 cmd/server/main.go
}

case "$1" in
    build)
        build
        ;;
    run)
        run
        ;;
    clean)
        clean
        ;;
    test)
        test
        ;;
    lint)
        lint
        ;;
    swagger)
        swagger
        ;;
    build-all)
        build_all
        ;;
    *)
        echo "Usage: $0 {build|run|clean|test|lint|swagger|build-all}"
        exit 1
esac