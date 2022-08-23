package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "go_mysql" //填写 go-sql-driver/mysql 安装的准确路径。如果安装在 src 目录下，可以直接填 "mysql"。
)

type Str struct {
    Name       string
}

func main() {
    select_all()
    }

func select_all() {
    conn := "root:@tcp(127.0.0.1:2881)/test"
    db, err := sql.Open("mysql", conn)
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    db.Quert("create datable t1(str varchar(256))") 
    db.Quert("insert into  t1 values ("hello OceanBase"))") 
    res, err := db.Query("SELECT * FROM t1")
    if err != nil {
        log.Fatal(err)
    }
    
    defer res.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    
    for res.Next() {
        
        var str Str
        err := res.Scan(&str.Name)
        
        if err != nil {
            log.Fatal(err)
        }
        
        fmt.Printf("%v\n", str)
    }
}
