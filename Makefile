.PHONY: build
build:
	rm -f staticweb
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o staticweb
