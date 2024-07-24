FROM golang:1.21

WORKDIR /app
COPY go.mod main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /whoami-server

FROM scratch
COPY --from=0 /whoami-server /whoami-server

ENTRYPOINT ["/whoami-server"]
