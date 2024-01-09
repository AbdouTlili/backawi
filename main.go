package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// product representation

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
	Discount float64
}

var Products []Product

func rootendpoint(rw http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : root")
	fmt.Fprintf(rw, "This is the root ")
}

func returnAllProducts(rw http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : return all products")
	json.NewEncoder(rw).Encode(Products)
}

func handleRequests() {
	http.HandleFunc("/", rootendpoint)
	http.HandleFunc("/products", returnAllProducts)
	http.ListenAndServe("localhost:8000", nil)
}

func main() {

	// init some dummy products
	Products = []Product{Product{Id: 1, Name: "Fromage", Quantity: 50, Price: 4.5, Discount: 0.0},
		Product{Id: 2, Name: "eggs", Quantity: 500, Price: 1.5, Discount: 20.0}}

	handleRequests()

}
