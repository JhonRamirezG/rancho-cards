package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//*This is a map of function that I can use in templates. This is to do things that are not inside the Go templates, like format Dates.
var functions = template.FuncMap{}

const filePathTemplates = "./templates/*.layout.html"

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//* Every Time I render a template I want to render it with the base layout + the unique content of the file.
	tc, err := CrateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	//* We take the template base on the tmpl that is the name of html, "ok" is to check if the template in tmpl exists if exist "ok" w'll have the value of true.
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	//* I need put the parse template that I have in memory into some bytes.
	buf := new(bytes.Buffer)
	//* Take the template that I have execute it, don't pass any data and store the value in buf variable.
	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
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
