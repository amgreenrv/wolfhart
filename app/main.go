package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Struct for inventory items
type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"prodname"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Img      string  `json:"img"`
	ProdDesc string  `json:"proddesc"`
	Featured int     `json:"featured"`
}

//Fetches FEATURED products from database
func getFeatured() []*Item {

	var password = os.Getenv("DB_PASSWORD")
	// Open up our database connection.
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(database:3306)/wolfhart", password))
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT prodname, img, featured FROM products WHERE featured = 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var items []*Item
	for results.Next() {
		var item Item
		err = results.Scan(&item.Name, &item.Img, &item.Featured)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, &item)
	}

	return items
}

//Fetches ALL products from database
func getInventory() []*Item {
	var password = os.Getenv("DB_PASSWORD")
	// Open up our database connection.
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(database:3306)/wolfhart", password))
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var items []*Item
	for results.Next() {
		var item Item
		err = results.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Img, &item.ProdDesc, &item.Featured)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, &item)
	}

	return items
}

///Featured Page Encode
func featuredPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) //Cors access to avoid error

	items := getFeatured() //calling the getFeatured function

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: featuredPage")
	json.NewEncoder(w).Encode(items)

}

//Product Page Encode
func productPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) //Cors access to avoid error

	items := getInventory() //calling the getInventory function

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: productsPage")
	json.NewEncoder(w).Encode(items)
}

//Main function with route handlers
func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/featured", featuredPage).Methods("GET")
	router.HandleFunc("/products", productPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

//CORS error fix - Cross-Origin Resources Sharing
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
