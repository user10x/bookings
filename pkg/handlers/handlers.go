package handlers

import (
	"fmt"
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"github.com/nickhalden/mynicceprogram/pkg/driver"
	"github.com/nickhalden/mynicceprogram/pkg/helpers"
	"github.com/nickhalden/mynicceprogram/pkg/models"
	"github.com/nickhalden/mynicceprogram/pkg/render"
	"github.com/nickhalden/mynicceprogram/repository"
	"github.com/nickhalden/mynicceprogram/repository/dbrepo"
	"net/http"
	"strconv"
	"time"
)

// Repo used by the Handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
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
	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Home returns home route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Calling Home route"))
	render.Template(w, "home.page.tmpl", &models.TemplateData{})

}

// SearchAvailability returns search results
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Calling Home route"))
	render.Template(w, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeRegistration(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Calling Home route"))
	render.Template(w, "make-registration.page.tmpl", &models.TemplateData{})
}

// PostRegistration returns search results
func (m *Repository) PostRegistration(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w,err)
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	layout := "2006-01-02"



	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w,err)
	}

	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
	}

	roomID, err := strconv.Atoi(r.FormValue("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	reservation := models.Reservation{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	err = m.DB.InsertReservation(reservation)

	if err != nil {
		helpers.ServerError(w, err)
	}

	w.Write([] byte("registration successful"))


}

// PostAvailability returns search results
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("start date is  %s %s", r.FormValue("start_date"), r.FormValue("end_date"))))
}


