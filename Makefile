installdir=/usr/local/bin

build:
	go build pre_commit.go

install: build
	sudo cp pre_commit $(installdir)/commit

