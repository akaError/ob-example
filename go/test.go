package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "mysql" //填写 go-sql-driver/mysql 安装的准确路径。如果安装在 src 目录下，可以直接填 "mysql"。
)

type City struct {
    Id         int
    Name       string
    Population int
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
    
    res, err := db.Query("SELECT * FROM cities")
    if err != nil {
        log.Fatal(err)
    }
    
    defer res.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    
    for res.Next() {
        
        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)
        
        if err != nil {
            log.Fatal(err)
        }
        
        fmt.Printf("%v\n", city)
    }
}
