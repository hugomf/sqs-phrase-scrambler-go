BINARY_NAME=main
FILE_NAME=main.go
DEPENDENCIES=word_scrambler.go word_assembler.go sqs_wrapper.go

build:
	go build -o bin/${BINARY_NAME} ${FILE_NAME} ${DEPENDENCIES}

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/${BINARY_NAME}-linux-arm ${FILE_NAME} ${DEPENDENCIES}
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 ${FILE_NAME} ${DEPENDENCIES}
	GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 ${FILE_NAME} ${DEPENDENCIES}
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-386 ${FILE_NAME} ${DEPENDENCIES}

.PHONY: clean
clean:
	go clean
	rm bin/${BINARY_NAME}
	rm bin/${BINARY_NAME}-linux-arm
	rm bin/${BINARY_NAME}-linux-arm64
	rm bin/${BINARY_NAME}-freebsd-386
	rm bin/${BINARY_NAME}-darwin-386


scrambler: build
	./bin/${BINARY_NAME} -s

assembler: build
	./bin/${BINARY_NAME} -a
