LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

deps:
	go mod tidy
	go mod download

build:
	$(LINUX_AMD64) go build -o esctl main.go

linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin

lint:
	golangci-lint run ./...

start: build
	go run main.go

dk-build:
	docker build -t esctl:latest .

mockgen:
	@go get go.uber.org/mock/mockgen@latest
	mockgen -source ./internal/file/editor.go -destination ./internal/file/editor_mock.go -package file
	mockgen -source ./internal/client/elastic.go -destination ./internal/client/elastic_mock.go -package client

test:
	go test ./... -covermode=count -count 1
