FROM golang

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/Crshi/go-spider
COPY . $GOPATH/src/github.com/Crshi/go-spider
RUN go build .

EXPOSE 8089
ENTRYPOINT ["./go-spider"]