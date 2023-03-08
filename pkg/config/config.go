package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
)

// AppConfig stores global config for the application to use
type AppConfig struct {
	InProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	Session      *scs.SessionManager
}
