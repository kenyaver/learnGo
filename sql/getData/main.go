package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	id int;
	model string;
	company string;
	price int;
}

func main() {
	db, err := sql.Open("mysql", "root:277353@/productdb");
	if err != nil {
		panic(err);
	}
	defer db.Close();

	row, err := db.Query("SELECT * FROM productdb.product");
	if err != nil {
		panic(err);
	}
	defer row.Close();

	productDB := []product{};
	for row.Next() {
		p := product{};
		err := row.Scan(&p.id, &p.model, &p.company, &p.price);
		if err != nil {
			fmt.Println(err);
			continue;
		}
		productDB = append(productDB, p);
	}

	for _, p := range productDB{
        fmt.Println(p.id, p.model, p.company, p.price)
    }
}