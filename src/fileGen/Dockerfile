FROM golang:latest

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 6969

CMD [ "/docker-gs-ping" ]