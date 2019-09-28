package dbhelp

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//SQLTable is a struct for SQLite tables
type SQLTable struct {
	Name      string
	Columns   []*SQLColumn
	IsCreated bool
}

//Create creates the table in the SQL database file
func (t *SQLTable) Create() error {
	if t.IsCreated {
		return nil
	}

	log.WithFields(log.Fields{"table": t}).Debugf("creating table %s\n", t.Name)

	var columns Query

	for _, col := range t.Columns {
		if col.IsPK {
			columns += Queryf(queryCreateTablePKColumPattern, col.Name, col.Type)
			continue
		}

		nullableString := ""
		if !col.Nullable {
			nullableString = "NOT NULL"
		}

		columns += Queryf(queryCreateTableColumPattern, col.Name, col.Type, nullableString)
	}

	q := Queryf(queryCreateTablePattern, t.Name, columns)

	err := q.Exec()
	if err != nil {
		log.WithError(err).Errorf("error when try create table %s", t.Name)
		return err
	}

	t.IsCreated = true
	log.Debugf("table %s is created\n", t.Name)
	return nil
}

//ExecSelectCols execute a SELECT query with specified columns
func (t *SQLTable) ExecSelectCols(columns []string, conditions []*ConditionGroup) (*sql.Rows, error) {
	var whereStr string
	if len(conditions) > 0 {
		whereStr = " WHERE "
		for _, cond := range conditions {
			whereStr += cond.ToString()
		}
	}
	cols := "*"
	if len(columns) > 0 {
		cols = ""
		for i, col := range columns {
			cols += fmt.Sprintf("[%s]", col)

			if i+1 < len(columns) {
				cols += ","
			}
		}
	}

	query := GetQuerySelect(t.Name, cols, whereStr)

	return query.ExecQuery()
}

//ExecSelectAllCols execute a `SELECT *...` query
func (t *SQLTable) ExecSelectAllCols(conditions []*ConditionGroup) (*sql.Rows, error) {
	return t.ExecSelectCols([]string{}, conditions)
}

//Insert inserts values in the table
func (t *SQLTable) Insert(values ...interface{}) error {
	if len(t.Columns)-1 != len(values) {
		return fmt.Errorf("there are differences between the columns and its values")
	}

	cols := ""
	vals := ""

	for i, col := range t.Columns {
		if col.IsPK {
			continue
		}

		cols += fmt.Sprintf("[%s]", col.Name)

		if col.Type == SQLTypeText && values[i] != nil {
			vals += fmt.Sprintf("'%v'", values[i])
		} else if values[i] == nil {
			vals += "NULL"
		} else {
			vals += fmt.Sprintf("%v", values[i])
		}

		if (i + 1) < len(t.Columns) {
			cols += ","
			vals += ","
		}
	}
	query := GetQueryInsert(t.Name, cols, vals)

	return query.Exec()
}

//Update updates the specified columns
func (t *SQLTable) Update(updateValues map[string]interface{}, conditions []*ConditionGroup) error {

	qUpdateSec := "SET "
	for colName, value := range updateValues {
		column := t.GetColByName(colName)

		if column.Type == SQLTypeText {
			qUpdateSec += fmt.Sprintf("[%s] = '%v,'", colName, value)
		} else {
			qUpdateSec += fmt.Sprintf("[%s] = %v,", colName, value)
		}
	}

	var whereStr string
	if len(conditions) > 0 {
		whereStr = " WHERE "
		for _, cond := range conditions {
			whereStr += cond.ToString()
		}
	}

	query := GetUpdateQuery(t.Name, qUpdateSec, whereStr)

	return query.Exec()
}

//GetColByName returns the column searching by name
func (t *SQLTable) GetColByName(name string) *SQLColumn {
	for _, col := range t.Columns {
		if col.Name == name {
			return col
		}
	}
	return nil
}
