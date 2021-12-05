package handlers

import (
	"net/http"

	"github.com/jhonrmz/rancho-cards/pkg/config"
	"github.com/jhonrmz/rancho-cards/pkg/models"
	"github.com/jhonrmz/rancho-cards/pkg/render"
)

//* Repo the repository used by the handlers.
var Repo *Repository

//* Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//* creates a new repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//* NewHandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//* Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	//!TemplateData is necessary because in the parameters of RenderTemplate is necessary to pass a TemplateData type and adding &{} its and empty TemplateData.
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//* About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//!TemplateData is necessary because in the parameters of RenderTemplate is necessary to pass a TemplateData type and adding &{} its and empty TemplateData.
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//* Shop is the shopping page handler
func (m *Repository) Shop(w http.ResponseWriter, r *http.Request) {
	//!TemplateData is necessary because in the parameters of RenderTemplate is necessary to pass a TemplateData type and adding &{} its and empty TemplateData.
	render.RenderTemplate(w, "shop.page.html", &models.TemplateData{})
}

//* Offers is the offers page handler
func (m *Repository) Offers(w http.ResponseWriter, r *http.Request) {
	//!TemplateData is necessary because in the parameters of RenderTemplate is necessary to pass a TemplateData type and adding &{} its and empty TemplateData.
	render.RenderTemplate(w, "offers.page.html", &models.TemplateData{})
}

//* Orders is the order page handler
func (m *Repository) Orders(w http.ResponseWriter, r *http.Request) {
	//!TemplateData is necessary because in the parameters of RenderTemplate is necessary to pass a TemplateData type and adding &{} its and empty TemplateData.
	render.RenderTemplate(w, "orders.page.html", &models.TemplateData{})
}
