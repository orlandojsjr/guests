package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var db *sql.DB

func InitDb() {
    var err error
    db, err := sql.Open("mysql", "root:abc123@tcp(127.0.0.1:3306)/guests")

    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}

