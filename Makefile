build:
	go build -o bin/WebServerGolang.exe cmd/web/main.go

clean:
	rm .\bin\WebServerGolang.exe