package dbhelp

//SQLRelOperator relational operator for SQL Conditions
type SQLRelOperator string

const (
	//Eq is Equals
	Eq SQLRelOperator = "=="
	//NEq is not Equals
	NEq SQLRelOperator = "<>"
	//Gt is Greather than
	Gt SQLRelOperator = ">"
	//Lt is Less than
	Lt SQLRelOperator = "<"
	//GtE is Greather than or equals
	GtE SQLRelOperator = ">="
	//LtE is Less than or equals
	LtE SQLRelOperator = "<="
	//StartW is Starts With
	StartW SQLRelOperator = "like '%v%%'"
	//EndW is Ends With
	EndW SQLRelOperator = "like '%%%v'"
	//Conts is Contains
	Conts SQLRelOperator = "like '%%%v%%'"
)
