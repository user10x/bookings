package helpers

import (
	"fmt"
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers setups app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println(status)
	http.Error(w, http.StatusText(status), status)pkg/config/config.go
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
