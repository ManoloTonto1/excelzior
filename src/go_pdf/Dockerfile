FROM golang:1.19.0-alpine3.16
RUN apk update
RUN apk upgrade 
RUN apk add libreoffice
RUN apk add openjdk8-jre

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 6970

CMD ["/docker-gs-ping"]