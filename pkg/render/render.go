package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jhonrmz/rancho-cards/pkg/config"
	"github.com/jhonrmz/rancho-cards/pkg/models"
)

//*This is a map of function that I can use in templates. This is to do things that are not inside the Go templates, like format Dates.
var functions = template.FuncMap{}

var app *config.AppConfig

//* NewTemplate sets the config for the template package.
func NewTemplates(a *config.AppConfig) {
	app = a
}

const filePathTemplates = "./templates/*.layout.html"

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//Set the variable tc here to can be used out side the if statement.
	var tc map[string]*template.Template
	//* This is to validate if we are in developer if app.UseCache is True otherwise we'll read the tempalates from the disk.
	if app.UseCache {
		//* Every Time I render a template I want to render it with the base layout + the unique content of the file.
		tc = app.TemplateCache
	} else {
		tc, _ = CrateTemplateCache()
	}
	//* We take the template base on the tmpl that is the name of html, "ok" is to check if the template in tmpl exists if exist "ok" w'll have the value of true.
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from templateCache")
	}
	//* I need put the parse template that I have in memory into some bytes.
	buf := new(bytes.Buffer)

	//* This is the call to set more data when is needed.
	td = AddDefaultData(td)

	//* Take the template that I have execute it, don't pass any data and store the value in buf variable.
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writting template to the browser", err)
	}
}

//* CreateTemplateCache creates a template cache as a map.
func CrateTemplateCache() (map[string]*template.Template, error) {
	//* Here I get the name of the file as index and a fully usable template as the value.
	myCache := map[string]*template.Template{}

	//* Bring all of the files from a specific location.

	//* This says go to the template folder and find everything end with .page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//* here is saying for every pages.html from the above call I'll get the index(i) and the page it's self(page).
	// Example as first iteration is:
	// 0:./templates/aboout.pages.html
	for _, page := range pages {
		//* If we only use the page we w'll get the full path and not just the name, we use the filepath.Base() to get the actual name of the file as:
		//about.page.html
		name := filepath.Base(page)

		//* We create a new template set base on the name of the file with and empty function
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//* Next we need to know if any of this tamplates match with any layouts. Match if any of this tamplates has a layout to match
		// Find any file inside the route templates with the end of layout.html
		matches, err := filepath.Glob(filePathTemplates)
		if err != nil {
			return myCache, err
		}

		//*If in the last lines(49) if there is at least 1 match of a file ending with layout.html the len w'll be > 0.
		if len(matches) > 0 {
			//* If I find the file with layout.html I want to parse that file.
			ts, err = ts.ParseGlob(filePathTemplates)
			if err != nil {
				return myCache, err
			}
		}

		//*Now we need to take the template set created and add it to the cache
		//The name is the name of the file like about.page.html and ts is te template set we created
		myCache[name] = ts
	}
	return myCache, nil
}
