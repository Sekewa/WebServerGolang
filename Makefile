build:
	go build -o bin/WebServerGolang.exe cmd/web/main.go

clean:
	del .\bin\WebServerGolang.exe