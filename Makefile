.PHONY: all say_hello generate clean
all: say_hello generate

say_hello:
	@echo "HELLO WORLD"

generate:
	@echo "Creatning empty text files..."
	touch file-{1..10}.txt

clean:
	@echo "Cleaning up..."
	rm *.txt

build_windows:
	GOOS=windows
	go build -ldflags -H=windowsgui -o builds/pomodoro-$$(echo -n $$(git describe --tags)).exe .

build_linux:
	GOOS=linux
	BUILD=pomodoro-$GOOS-$$(echo -n $$(git describe --tags))
	go build -o builds/$BUILD .
	tar -czf releases/$BUILD.tgz builds/*
	rm builds/*

build_mac:
	#dmg
	GOOS=darwin