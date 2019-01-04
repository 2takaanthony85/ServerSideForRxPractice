FROM golang:1.9

RUN mkdir /go/src/serversiderxswift \
    && go get -u github.com/gin-gonic/gin \
    && go get -u github.com/go-sql-driver/mysql \
    && go get -u github.com/go-xorm/xorm \
    && go get -u github.com/wakashiyo/serverSideForRxPractice/todos \
    && go get -u github.com/wakashiyo/serverSideForRxPractice/users
    
COPY main.go /go/src/serversiderxswift

CMD ["go", "run", "/go/src/serversiderxswift/main.go"]