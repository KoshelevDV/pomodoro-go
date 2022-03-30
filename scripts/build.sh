go build -ldflags -H=windowsgui -o builds/pomodoro-$(echo -n $(git describe --tags)).exe .
