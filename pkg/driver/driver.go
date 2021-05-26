package driver

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

//DB holds the connection for database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConnections = 10
const maxIdleConnections = 5
const maxDbConnectionLifetime = 5 * time.Minute

// ConnectSQL creates a database pool for postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err) // does not start up the app if the db is not up
	}
	d.SetMaxOpenConns(maxOpenDbConnections)
	d.SetMaxIdleConns(maxIdleConnections)
	d.SetConnMaxLifetime(maxDbConnectionLifetime)
	dbConn.SQL = d
	return dbConn, nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
