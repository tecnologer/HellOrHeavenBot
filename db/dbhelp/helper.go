package dbhelp

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

//Connection instance of connection for the database
var Connection *sql.DB

//Open creates a connection for database
func Open() (err error) {
	Connection, err = sql.Open("sqlite3", "./db/dbFiles/luciack.db")

	if err != nil {
		log.Debug("connection to database is created")
	}
	return
}

//Close closes the connection for database
func Close() {
	Connection.Close()
	log.Debug("connection to database is closed")
}

//BeginTran starts a new transaction
func BeginTran() (*sql.Tx, error) {
	if Connection == nil {
		return nil, fmt.Errorf("connection is closed")
	}

	return Connection.Begin()
}
