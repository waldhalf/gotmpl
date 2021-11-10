package handlers

import (
	"net/http"

	"github.com/waldhalf/gotmpl/pkg/config"
	"github.com/waldhalf/gotmpl/pkg/models"
	"github.com/waldhalf/gotmpl/pkg/render"
)

// Get App Config
// var app = config.GetConfig()

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repositoty
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

