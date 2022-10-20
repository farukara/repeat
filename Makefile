.PHONY: all
all: re

re: main.go
	go build -o $@ $^
	cp $@ /usr/local/bin/
	rm re
