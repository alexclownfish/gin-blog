FROM golang:1.17.7
MAINTAINER Alex “ywz0207@163.com”
WORKDIR /opt/web-system/go-build
ADD . $GOPATH/src/github.com/gin-blog
RUN go build .
EXPOSE 6064
ENTRYPOINT ["./gin-blog"]