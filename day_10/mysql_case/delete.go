package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}

func main() {

	_, err := Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("delete succ")
}
