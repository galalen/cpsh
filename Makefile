BINARY_NAME=cpsh

build:
	go build -o ${BINARY_NAME} cmd/cpsh.go

build-mac:
	GOOS=darwin go build -o ${BINARY_NAME} cmd/cpsh.go

build-linux:
	GOOS=linux go build -o ${BINARY_NAME} cmd/cpsh.go

build-windows:
	GOOS=windows go build -o ${BINARY_NAME} cmd/cpsh.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
