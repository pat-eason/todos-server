MAIN_PACKAGE=./cmd/api.go
BINARY_NAME=todo-server
BINARY_PATH=./bin

build:
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_PATH}/${BINARY_NAME}-darwin ${MAIN_PACKAGE}
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_PATH}/${BINARY_NAME}-linux ${MAIN_PACKAGE}

run-darwin: build
	${BINARY_PATH}/${BINARY_NAME}-darwin

run-linux: build
	${BINARY_PATH}/${BINARY_NAME}-linux

clean:
	go clean
	rm ${BINARY_PATH}/${BINARY_NAME}-darwin
	rm ${BINARY_PATH}/${BINARY_NAME}-linux
