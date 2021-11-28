package main

import (
	"fmt"
	"net/http"

	"github.com/jhonrmz/rancho-cards/pkg/handlers"
)

const portNumber = ":8080"

//* main is the main func
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/shop", handlers.Shop)
	http.HandleFunc("/offers", handlers.Offers)
	http.HandleFunc("/orders", handlers.Orders)

	//Start the web server.
	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
