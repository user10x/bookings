package handlers

import (
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"github.com/nickhalden/mynicceprogram/pkg/models"
	"github.com/nickhalden/mynicceprogram/pkg/render"
	"net/http"
)

// Repo used by the Handlers
var Repo *Repository

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

//Health check if the sample route is serving
func (m *Repository) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Calling Health route"))
}

// About returns about content
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Calling About route"))
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello from template"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Home returns home route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Calling Home route"))
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}
