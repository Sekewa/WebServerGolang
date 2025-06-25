build:
	go build -o out/WebServerGolang.exe cmd/web/main.go

clean:
	del .\out\WebServerGolang.exe