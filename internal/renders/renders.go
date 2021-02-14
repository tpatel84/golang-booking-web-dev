package renders

import (
	"bytes"
	"github.com/justinas/nosurf"
	"github.com/tpatel84/golang-booking-web-dev/internal/config"
	"github.com/tpatel84/golang-booking-web-dev/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		// Create a new template cache
		tc, _ = CreateTemplateCache()
	}

	// Following call will create a template cache every time
	//tc, err := CreateTemplateCache()
	//if err != nil {
	//	log.Fatal(err)
	//}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}

	//parsedTpl, _ := template.ParseFiles("templates/" + tmpl)
	//err = parsedTpl.Execute(w, nil)
	//if err != nil {
	//	log.Println("Error parsing template", err)
	//	return
	//}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("templates/*.page.tmpl")
	if err != nil {
		log.Println("Failed to find templates files")
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Failed to create templates")
			return myCache, err
		}
		matches, err := filepath.Glob("templates/*.layout.tmpl")
		if err != nil {
			log.Println("Failed to get matching templates")
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
