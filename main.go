package main

import (
	"database/sql"
	"net/http"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func main() {
	pruduct := Product {
		ID: "1",
		Name: "Product 1",
		Price: 100.00,
	}
	err := SaveProduct(pruduct)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func SaveProduct(product Product) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()
	
	stmt, err != db.Prepare("INSERT INTO Products(id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}