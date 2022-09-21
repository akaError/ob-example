gopath=$(cd `dirname $0`;pwd)
mkdir src
echo "export GOPATH="$gopath"" >> ~/.bash_profile
source ~/.profile
export GOPROXY=https://goproxy.io&&export GO111MODULE=on
go get github.com/go-sql-driver/mysql
