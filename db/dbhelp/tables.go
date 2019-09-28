package dbhelp

import log "github.com/sirupsen/logrus"

//SQLTable is a struct for SQLite tables
type SQLTable struct {
	Name      string
	Columns   []*SQLColumn
	IsCreated bool
}

//Create creates the table in the SQL database file
func (t *SQLTable) Create() error {
	log.WithFields(log.Fields{"table": t}).Debugf("creating table %s\n", t.Name)
	if t.IsCreated {
		return nil
	}

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
