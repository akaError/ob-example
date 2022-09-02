gopath=$(cd `dirname $0`;pwd)
go env -w GO111MODULE=auto
echo "export GOPATH="$gopath >> ~/.profile
source ~/.profile
go get github.com/go-sql-driver/mysql
