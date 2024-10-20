BINARY=/usr/bin/soup
SOURCE=soup.go

install:
	go build -o ${BINARY} ${SOURCE}

clean:
	go clean
	rm ${BINARY}
