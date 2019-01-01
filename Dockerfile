FROM golang:1.9

RUN mkdir /go/src/serversiderxswift \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/wakashiyo/serverSideForRxPractice/todos
COPY main.go /go/src/serversiderxswift

CMD ["go", "run", "/go/src/serversiderxswift/main.go"]