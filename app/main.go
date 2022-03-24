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

//Struct for inventory items, provides structure for the fetched data
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
//[]*Item is a slice of pointers to the struct Item, this allows me to use &item reference below in the append.
//Also known as dereferencing or indirecting
func getFeatured() []*Item {
	//Used to create secure variables for database
	var password = os.Getenv("DB_PASSWORD")
	var database = os.Getenv("DB_NAME")

	//Open the database connection using the secure variables - %s is basic string printing
	//username, password, port, database name
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(database:3306)/%s", password, database))
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error()) // error handling
	}
	defer db.Close() //defer tells the database not to close until the surrounding function is complete.

	// Execute the query
	//Query only selects the data we need to serve the featured product request/route
	results, err := db.Query("SELECT prodname, img, featured FROM products WHERE featured = 1")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error()) // error handling
	}
	//Scns through the query above, .Next directs it to the next item on the list then appends the items to
	var items []*Item
	for results.Next() {
		var item Item
		err = results.Scan(&item.Name, &item.Img, &item.Featured)
		if err != nil {
			panic(err.Error()) // error handling
		}
		items = append(items, &item)
	}

	return items //returning the newly appended slice of struct with data.
}

//Fetches ALL products from database
func getInventory() []*Item {
	//Used to create secure variables for database
	var password = os.Getenv("DB_PASSWORD")
	var database = os.Getenv("DB_NAME")

	//Open the database connection using the secure variables
	//username, password, port, database name
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(database:3306)/%s", password, database))
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error()) // error handling
	}
	defer db.Close() //defer tells the database not to close until the surrounding function is complete.

	// Execute the query, select everything from the products table
	results, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error()) // error handling
	}

	var items []*Item
	for results.Next() {
		var item Item
		err = results.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Img, &item.ProdDesc, &item.Featured)
		if err != nil {
			panic(err.Error()) // error handling
		}
		items = append(items, &item)
	}

	return items
}

///Featured Page Encode
func featuredPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) //Cors access to avoid error

	items := getFeatured() //calling the getFeatured function

	//returns the 200 (StatusOK) status code
	w.WriteHeader(http.StatusOK)
	//header establishes the type of content the browser expects to receive (json)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: featuredPage") //console log to determine if this endpoint was hit successfully
	json.NewEncoder(w).Encode(items)          //encodes the response in json

}

//Product Page Encode
func productPage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w) //Cors access to avoid error

	items := getInventory() //calling the getInventory function

	w.WriteHeader(http.StatusOK)
	//header establishes the type of content the browser expects to receive (json)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: productsPage") //console log to determine if this endpoint was hit successfully
	json.NewEncoder(w).Encode(items)
}

//Main function with route handlers for products and featured end points
func main() {
	router := mux.NewRouter().StrictSlash(true)
	//StrictSlash is a parameter on Mux router that tells it to redirect URL routes with trailing slashes /featured/ vs /featured

	router.HandleFunc("/featured", featuredPage).Methods("GET")
	router.HandleFunc("/products", productPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
	//logging any errors, open and listening to port 8080

}

//CORS error fix - Cross-Origin Resources Sharing, * = share with all
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
