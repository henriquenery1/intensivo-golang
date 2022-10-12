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
	println(pruduct.ID)

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func SaveProduct(product Product) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
}