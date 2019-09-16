package db

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

type query string

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

func execQueryNoResult(query query) (err error) {
	if Connection == nil {
		err = Open()

		if err != nil {
			return
		}
	}

	_, err = Connection.Exec(query.String())

	log.WithFields(log.Fields{
		"query":     query,
		"hasResutl": false,
	}).Debug("new query executed")
	return
}

func execQuery(q query) (*sql.Rows, error) {
	log.WithFields(log.Fields{
		"query":     q,
		"hasResutl": true,
	}).Debug("new query executed")

	return Connection.Query(q.String())
}

func (q query) String() string {
	return string(q)
}

//queryf creates a query using a format
func queryf(pattern string, values ...interface{}) query {
	return query(fmt.Sprintf(pattern, values...))
}
