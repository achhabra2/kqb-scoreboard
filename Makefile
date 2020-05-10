compile:
	echo "Compiling for mac and windows"
	GOOS=windows GOARCH=amd64 go build -o kqb-scoreboard.exe .
	GOOS=darwin GOARCH=amd64 go build -o kqb-scoreboard .