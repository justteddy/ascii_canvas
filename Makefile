test:
	go test -cover -v -parallel 8 ./...
dep:
	go mod tidy
lint:
	golangci-lint -v run