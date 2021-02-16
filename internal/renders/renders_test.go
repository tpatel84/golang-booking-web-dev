package renders

import (
	"github.com/tpatel84/golang-booking-web-dev/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	// http.request requires session object
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	// Add key and value into session for test
	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplate = "./../../templates"

	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplate(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Rendered template doesn't exist")
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplate = "./../../templates"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	// Create dummy GET request
	r, err := http.NewRequest("GET", "/some_url", nil)
	if err != nil {
		return nil, err
	}
	// Get context from http.request
	ctx := r.Context()
	// Get session data for given token
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// Put context back into request
	r = r.WithContext(ctx)

	return r, nil
}

