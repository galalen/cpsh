BINARY_NAME=cpsh

build:
	GOARCH=amd64 go build -o ${BINARY_NAME} cmd/cpsh.go

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} cmd/cpsh.go

build-linux:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} cmd/cpsh.go

build-windows:
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME} cmd/cpsh.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
