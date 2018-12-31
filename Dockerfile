FROM golang:1.9

RUN mkdir /go/src/serversiderxswift
COPY main.go /go/src/serversiderxswift

CMD ["go", "run", "/go/src/serversiderxswift/main.go"]