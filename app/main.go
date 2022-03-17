package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//struct for inventory items
type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"prodname"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Img      string  `json:"img"`
	ProdDesc string  `json:"proddesc"`
	Featured string  `json:"featured"`
}

//Fetches all inventory
func getInventory() []*Item {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:fallon87@tcp(database:3306)/wolfhart")
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

//Fetches just the featured inventory
func getFeatured() []*Item {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:fallon87@tcp(database:3306)/wolfhart")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM products WHERE featured=`true`")
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

func homePage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	items := getFeatured()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(items)

}

func productPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	items := getInventory()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: productsPage")
	json.NewEncoder(w).Encode(items)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/products", productPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
