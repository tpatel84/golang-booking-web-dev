package handlers

import (
	"github.com/tpatel84/golang-booking-web-dev/pkg/config"
	"github.com/tpatel84/golang-booking-web-dev/pkg/models"
	"github.com/tpatel84/golang-booking-web-dev/pkg/renders"
	"net/http"
)


// Repo the repository used the handlers
var Repo *Repository

// Repository is the type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
