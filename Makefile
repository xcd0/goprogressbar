

build:
	GOOS=windows GOARCH=amd64 go build -o goprogressbar_win
	GOOS=darwin  GOARCH=amd64 go build -o goprogressbar_mac
	GOOS=linux   GOARCH=amd64 go build -o goprogressbar_linux

run: build
	./goprogress 500

clean:
	rm goprogressbar_win
	rm goprogressbar_mac
	rm goprogressbar_linux
