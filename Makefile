.PHONY: server clean

SERVER := file_transfer
BUILD_CMD := GOOS=linux GOARCH=arm64 go build

server:
	${BUILD_CMD} -o ${SERVER}
clean:
	@rm -r ${SERVER} ${SCRIPT}