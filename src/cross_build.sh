DIR=$(cd ../; pwd)
export GOPATH=$DIR:$GOPATH
GOOS=linux GOARCH=amd64 go build qufop.go
