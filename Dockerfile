FROM golang:1.21

WORKDIR /app
COPY go.mod main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /whoami-server

ENTRYPOINT ["/whoami-server"]
