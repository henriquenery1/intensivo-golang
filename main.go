package main

import "net/http"

type Product struct {
	ID string
	Name string
	Price float64
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}