package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func main() {

	fmt.Println("Hello World")

	Products = []Product{
		{"1", "Chair", 100, 100.00},
		{"2", "Desk", 200, 200.00},
	}

	handleRequests()
}

func handleRequests() {

	http.HandleFunc("/", homepage)

	http.HandleFunc("/products", returnAllProducts)

	http.HandleFunc("/product/", getProduct)

	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the new World of the GO.")

	log.Println("EndPoint hit: homepage")

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {

	log.Println("EndPoint hit : ReturnaAllProducts")

	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL.Path)

	key := r.URL.Path[len("/product/"):]

	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
