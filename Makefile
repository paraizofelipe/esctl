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
