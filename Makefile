installdir=/usr/local/bin

build:
	go build main.go

install: build
	sudo cp main $(installdir)/commit

