package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	urlConnection := "root@/sas?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", urlConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexao está aberta")

	rows, err := db.Query("SELECT * FROM properties")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)
}
