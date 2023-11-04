package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:277353@/productdb");

	if err != nil {
		panic(err);
	}
	defer db.Close();

	result, err := db.Exec("insert into productdb.product (model, company, price) values (?, ?, ?)", "iPhone X", "Apple", 72000);
	if err != nil {
		panic(err);
	}

	fmt.Println(result.LastInsertId());
	fmt.Println(result.RowsAffected());
}