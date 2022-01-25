FROM golang:1.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY start.sh ./

RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 GO111MODULE=on go build --ldflags='-extldflags=-static' -o ./hashapi

EXPOSE 9000

CMD [ "./start.sh" ]