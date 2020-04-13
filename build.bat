SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -a -installsuffix cgo -o go-spider .
docker build -t go-spider .

exit