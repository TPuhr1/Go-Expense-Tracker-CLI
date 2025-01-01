# Makefile for building and running expense-tracker

build:
	go build -o expense-tracker expense-tracker.go

run:
	./expense-tracker $(args)

install:
	sudo mv expense-tracker /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/expense-tracker
