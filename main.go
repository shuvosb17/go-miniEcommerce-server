package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImgUrl      string `json:"ImageUrl"`
}

var productList []Product // To store all the product in array slice

func helloHandler(w http.ResponseWriter, r *http.Request) { // Hello route 1
	fmt.Fprintln(w, "Hello, This is Shuvo")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) { // aboute route 2
	fmt.Fprintln(w, "Shut the heckk!")
}
func getProducts(w http.ResponseWriter, r *http.Request) { // getProduct route 3

	HandleCors(w)         //Handle all cors
	handlePreflight(w, r) //Handle all preflight request

	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", 400)
		return
	}

	sendData(w, productList, 200)

}

func createProduct(w http.ResponseWriter, r *http.Request) { //createProduct route 4
	HandleCors(w)
	handlePreflight(w, r)

	if r.Method != "POST" {
		http.Error(w, "Please give valid Request", 400)
		return
	}

	var newProduct Product //create new product in fronted

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct) // converting json file

	if err != nil { // nil = correct
		fmt.Println(err)
		http.Error(w, "Please give the valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct) //appending newProduct into productList

	sendData(w, newProduct, 201)
}

func HandleCors(w http.ResponseWriter) { //CorsHandler function maintaing Single Responsibility principle (SRP)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Shuvo")
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
}

func handlePreflight(w http.ResponseWriter, r *http.Request) { //PreflightHandler function maintaing Single Responsibility principle (SRP)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}
func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-Product", createProduct)

	fmt.Println("Server running on port: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil { //pstv response = nil
		fmt.Println("Server error at", err)
	}
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Yellow",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-vector/watercolor-orange-background_52683-10330.jpg?semt=ais_hybrid&w=740&q=80",
	}

	prod2 := Product{
		ID:          2,
		Title:       "Orange",
		Description: "Yellow",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-vector/watercolor-orange-background_52683-10330.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Orange",
		Description: "Yellow",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-vector/watercolor-orange-background_52683-10330.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prod4 := Product{
		ID:          4,
		Title:       "Orange",
		Description: "Yellow",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-vector/watercolor-orange-background_52683-10330.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prod5 := Product{
		ID:          5,
		Title:       "Orange",
		Description: "Yellow",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-vector/watercolor-orange-background_52683-10330.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prod6 := Product{
		ID:          6,
		Title:       "Orange",
		Description: "Yellow",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-vector/watercolor-orange-background_52683-10330.jpg?semt=ais_hybrid&w=740&q=80",
	}

	productList = append(productList, prod1)
	productList = append(productList, prod2)
	productList = append(productList, prod3)
	productList = append(productList, prod4)
	productList = append(productList, prod5)
	productList = append(productList, prod6)
}
