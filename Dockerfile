FROM golang:1.14.7
MAINTAINER GitHub, Inc.

WORKDIR /go/src/github.com/CHENXCHEN/lfs-server-go

COPY . .

RUN go build

EXPOSE 8080

CMD /go/src/github.com/CHENXCHEN/lfs-server-go/lfs-server-go
