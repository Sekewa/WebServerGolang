FROM golang:1.24.4

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/WebServerGolang cmd/web/main.go

# if you want to change port for the container expose its here
EXPOSE 8080

ENTRYPOINT ["bin/WebServerGolang.exe"]
# and if you want to change the port for WebServerGolang exposure in container
# its here
CMD [ "--p","8080" ]