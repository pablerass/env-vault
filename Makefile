export GO111MODULE=on
VERSION=$(shell git describe --tags --candidates=1 --dirty)
FLAGS=-X main.Version=$(VERSION) -s -w
#CERT="Developer ID Application: 99designs Inc (NRM9HVJ62Z)"
SRC=$(shell find . -name '*.go')

.PHONY: all clean release

all: env-vault-linux-amd64 env-vault-windows-386.exe env-vault-freebsd-amd64 #env-vault-darwin-amd64

clean:
	rm -f env-vault-linux-amd64 env-vault-darwin-amd64 env-vault-windows-386.exe env-vault-freebsd-amd64

#release: all
#	codesign -s $(CERT) env-vault-darwin-amd64
#	@echo "\nTo update homebrew-cask run\n\n    cask-repair -v $(shell echo $(VERSION) | sed 's/v\(.*\)/\1/') env-vault\n"

env-vault-linux-amd64: $(SRC)
	GOOS=linux GOARCH=amd64 go build -o $@ -ldflags="$(FLAGS)" .

env-vault-darwin-amd64: $(SRC)
	GOOS=darwin GOARCH=amd64 go build -o $@ -ldflags="$(FLAGS)" .

env-vault-windows-386.exe: $(SRC)
	GOOS=windows GOARCH=386 go build -o $@ -ldflags="$(FLAGS)" .

env-vault-freebsd-amd64: $(SRC)
	GOOS=freebsd GOARCH=amd64 go build -o $@ -ldflags="$(FLAGS)" .