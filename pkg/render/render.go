package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a template
// TODO: remove func
func RenderTemplate2(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var html *template.Template
	var err error

	// check we are already the template
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("create the template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("use cache the template")
	}

	html = tc[t]
	err = html.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.html",
	}
	// parsing the template
	html, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = html

	return nil
}
