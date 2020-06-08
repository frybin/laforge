## TESTING REDIS WITH GO ##
FROM golang:1.14

WORKDIR /go/src/app
COPY . .
COPY ./laforge_proto /go/src/github.com/frybin/laforge/grpc-alpha/laforge_proto

RUN go get -d -v ./...
RUN go build ./server.go

CMD ["./server"]