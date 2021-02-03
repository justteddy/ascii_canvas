test:
	go test -cover -v -parallel 8 ./...
dep:
	go mod tidy