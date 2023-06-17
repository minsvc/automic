FROM golang:1.20.5-alpine3.18

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /go

RUN apk update && apk add --no-cache git
RUN git clone https://github.com/minsvc/automic.git

WORKDIR /go/automic
RUN go build -v main.go
EXPOSE $PORT

ENTRYPOINT ["./main"]