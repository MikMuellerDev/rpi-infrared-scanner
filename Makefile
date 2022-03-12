gosetup:
	go mod tidy

build: gosetup
	GOARCH=arm go build -o scanner main.go