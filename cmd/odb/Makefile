BIN=odb

.PHONY:inst build-win build-mac-intel build-mac-arm
.SILENT: inst
inst:
	go build .
	sudo mv $(BIN) /usr/local/bin 

build-win:
	GOOS=windows GOARCH=amd64 go build -o odb-win .

build-mac-intel:
	GOOS=darwin GOARCH=amd64 go build -o odb-mac-amd .

build-mac-arm:
	GOOS=darwin GOARCH=arm64 go build -o odb-mac-arm .