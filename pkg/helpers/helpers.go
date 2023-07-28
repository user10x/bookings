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
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Printf("%v\n", debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
