package config

import "log"

// AppConfig stores global config for the application to use
type AppConfig struct {
	InProduction bool
	InfoLog      *log.Logger
}
