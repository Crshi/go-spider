FROM scratch

WORKDIR $GOPATH/src/github.com/Crshi/go-spider
COPY . $GOPATH/src/github.com/Crshi/go-spider

EXPOSE 8089
ENTRYPOINT ["./go-spider"]