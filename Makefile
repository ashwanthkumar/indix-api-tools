VERSION = 0.1.0-dev
BASE_PACKAGE = github.com/ashwanthkumar/indix-api-tools

setup:
	glide install

build-json2csv:
	go build -o json2csv.out ${BASE_PACKAGE}/json2csv

build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o json2csv-linux-amd64 ${BASE_PACKAGE}/json2csv
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o json2csv-darwin-amd64 ${BASE_PACKAGE}/json2csv
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -X main.Version=${VERSION}" -v -o json2csv-win64.exe ${BASE_PACKAGE}/json2csv

install: build
	sudo install -d /usr/local/bin
	sudo install -c ${APPNAME} /usr/local/bin/${APPNAME}

uninstall:
	sudo rm /usr/local/bin/${APPNAME}
