package dbrepo

import (
	"database/sql"
	"github.com/nickhalden/mynicceprogram/pkg/config"
	"github.com/nickhalden/mynicceprogram/repository"
)

//
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// mysql to add another database

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
