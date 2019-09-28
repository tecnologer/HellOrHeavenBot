package dbhelp

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Query string of query
type Query string

const (
	queryCreateTablePattern        = "CREATE TABLE IF NOT EXISTS [%s] (%s);"
	queryCreateTablePKColumPattern = "[%s] %s NOT NULL primary key AUTOINCREMENT"
	queryCreateTableColumPattern   = "[%s] %s %s" //[<column-name>] <column-type> [not null]
)

//ExecQuery executes a query that returns rows, typically a SELECT.
func (q Query) ExecQuery() (*sql.Rows, error) {
	log.WithFields(log.Fields{
		"query":     q,
		"hasResutl": true,
	}).Debug("new query executed")

	return Connection.Query(q.String())
}

//Exec executes a query without returning any rows.
func (q Query) Exec() (err error) {
	if Connection == nil {
		err = Open()

		if err != nil {
			return
		}
	}

	_, err = Connection.Exec(q.String())

	log.WithFields(log.Fields{
		"query":     q,
		"hasResutl": false,
	}).Debug("new query executed")

	return
}

func (q Query) String() string {
	return string(q)
}

//Queryf creates a query using a format
func Queryf(pattern string, values ...interface{}) Query {
	return Query(fmt.Sprintf(pattern, values...))
}
