installdir=/usr/local/bin

build:
	go build git-dpcommit.go

install: build
	sudo cp git-dpcommit $(installdir)/git-dpcommit

