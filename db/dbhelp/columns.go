package dbhelp

//ColumnType type for SQL columns
type ColumnType string

//SQLColumn is the struct for columns table
type SQLColumn struct {
	Name     string
	Type     ColumnType
	Nullable bool
	IsPK     bool
}

const (
	//SQLTypeInt is a constant string for INTEGER
	SQLTypeInt ColumnType = "INTEGER"
	//SQLTypeText is a constant string for TEXT
	SQLTypeText ColumnType = "TEXT"
)

//NewIntCol creates an integer sql column not nullable
func NewIntCol(name string) *SQLColumn {
	return &SQLColumn{
		Name:     name,
		Type:     SQLTypeInt,
		Nullable: false,
		IsPK:     false,
	}
}

//NewIntNilCol creates an integer sql column nullable
func NewIntNilCol(name string) *SQLColumn {
	return &SQLColumn{
		Name:     name,
		Type:     SQLTypeInt,
		Nullable: true,
		IsPK:     false,
	}
}

//NewTextCol creates a text sql column not nullable
func NewTextCol(name string) *SQLColumn {
	return &SQLColumn{
		Name:     name,
		Type:     SQLTypeText,
		Nullable: false,
		IsPK:     false,
	}
}

//NewTextNilCol creates a text sql column nullable
func NewTextNilCol(name string) *SQLColumn {
	return &SQLColumn{
		Name:     name,
		Type:     SQLTypeText,
		Nullable: true,
		IsPK:     false,
	}
}

//NewPKCol creates a PK column, integer not null autoincrement
func NewPKCol(name string) *SQLColumn {
	col := NewIntCol(name)
	col.IsPK = true
	return col
}
