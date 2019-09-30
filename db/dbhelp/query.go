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
	queryCreateTablePKColumPattern = "[%s] %s NOT NULL primary key AUTOINCREMENT,"
	queryCreateTableColumPattern   = "[%s] %s %s" //[<column-name>] <column-type> [not null]
	querySelectCols                = "SELECT %s FROM [%s]%s;"
	queryInsert                    = "INSERT INTO [%s] (%s) VALUES (%s);"
	queryUpdate                    = "UPDATE [%s] SET %s%s;"
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

//GetQuerySelect returns a Query (SELECT) with the list of columns and conditions specified
func GetQuerySelect(tableName, cols, where string) Query {
	return Queryf(querySelectCols, cols, tableName, where)
}

//GetQueryInsert returns a Query (INSERT) to insert data in the table
func GetQueryInsert(tableName, cols, values string) Query {
	return Queryf(queryInsert, tableName, cols, values)
}

//GetUpdateQuery returns a Query (UPDATE) to update data in the table
func GetUpdateQuery(tableName, updateValues, where string) Query {
	return Queryf(queryUpdate, tableName, updateValues, where)
}
