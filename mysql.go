package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql","root:123bgn@tcp(192.168.105.112:3306)/msoa?charset=utf8")
	id := 1
	rows, err := db.Query("SELECT comment FROM comments WHERE user_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, id)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
