FROM golang:1.12.7-alpine3.10 AS builder

RUN apk update && apk add --no-cache git
RUN go get google.golang.org/grpc
RUN mkdir -p /go/src/grp
RUN mkdir -p /go/src/grp/service
RUN mkdir -p /go/src/grp/server
ADD ./main.go /go/src/grp/server
COPY ./service.pb.go /go/src/grp/service
WORKDIR /go/src/grp/server
EXPOSE 50005
RUN go build -o main .
CMD ["/go/src/grp/server/main"]