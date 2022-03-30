go build -ldflags -H=windowsgui -o builds/pomodoro-$(echo $(git describe --tags))-windows.exe .
