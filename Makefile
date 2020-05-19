compile:
	export CGO_ENABLED=1
	packr2
	echo "Compiling for mac and windows"
	GOOS=windows GOARCH=amd64 go build -o kqb-scoreboard.exe .
	GOOS=darwin GOARCH=amd64 go build -o kqb-scoreboard .
	packr2 clean
fyne-win:
	packr2
	fyne-cross windows
	packr2 clean

fyne-mac:
	packr2
	GOOS=darwin GOARCH=amd64 go build -o kqb-scoreboard .
	fyne package -os darwin -icon Icon.png
	packr2 clean