package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
if err != nil {
    panic(err)
}
defer db.Close()
fmt.Print("Подключено")
}