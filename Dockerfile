FROM golang:1.18-bullseye

RUN apt-get update
RUN apt-get install -y supervisor # Installing supervisord

RUN mkdir /src
WORkDIR /src

ADD . .
RUN go mod download

ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf 

RUN go build cmd/server/server.go
RUN go build cmd/worker/worker.go
EXPOSE 8080

ENTRYPOINT ["/usr/bin/supervisord"]
