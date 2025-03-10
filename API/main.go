package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)

	myRouter.HandleFunc("/products", returnAllProducts)

	myRouter.HandleFunc("/product/{id}", getProduct)

	http.ListenAndServe(":8080", myRouter)
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

	//log.Println(r)

	vars := mux.Vars(r)

	key := vars["id"]

	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}
