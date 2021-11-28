package handlers

import (
	"net/http"

	"github.com/jhonrmz/rancho-cards/pkg/render"
)

//* Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

//* About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

//* Shop is the shopping page handler
func Shop(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "shop.page.html")
}

//* Offers is the offers page handler
func Offers(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "offers.page.html")
}

//* Orders is the order page handler
func Orders(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "orders.page.html")
}
