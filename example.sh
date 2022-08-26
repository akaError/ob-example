#!/bin/bash

help()
{
    cat <<- EOF
    
    -h --help       show help

    -r --run        [python3 | go | java]
    
EOF
    exit 0
}
while [ -n "$1" ];do #这里通过判断$1是否存在
    case $1 in
    -r|--run) name=$2 # $2才是我们希望输出的参数
    shift 2;; # 将参数后移2个，进入下一个参数的判别
    -h|--help) help;; # function help is called
    --) shift;break;; # end of options
    -*) echo "error: no such option $1."; exit 1;;
    *) break;;
    esac
done

python="python3"
java="java"
go="go"
echo $name
if [ -n "$name" ];then
    echo 'start to connect OceanBase with '$name
    if [ $python == $name ]
    then
      python3 /workspace/ob-example/python/Test.py 
    elif [ "$java" == $name ]
    then
      cd /workspace/ob-example/java&&javac -cp ./mysql-connector-java-5.1.47.jar Test.java >> ~/.output&&java -cp .:mysql-connector-java-5.1.47.jar Test
    elif [ $go == $name ]
    then
      go run /workspace/ob-example/go/Test.go
    else
      echo 'Not Support Language: '$name
    fi
    echo 
fi
