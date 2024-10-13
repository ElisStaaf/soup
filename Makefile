BINARY=soup
SOURCE=soup.go

build:
	go build -o ${BINARY} ${SOURCE}

run:
	go build -o ${BINARY} ${SOURCE}
	./${BINARY}

clean:
	go clean
	rm ${BINARY}
