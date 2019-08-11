FROM golang:1.12.7-alpine3.10 AS builder

RUN apk update && apk add --no-cache git
RUN go get google.golang.org/grpc
RUN mkdir -p /go/src/grp
RUN mkdir -p /go/src/grp/service
RUN mkdir -p /go/src/grp/client
ADD ./main.go /go/src/grp/client
COPY ./service.pb.go /go/src/grp/service
WORKDIR /go/src/grp/client
RUN go build -o main .
CMD ["/go/src/grp/client/main"]