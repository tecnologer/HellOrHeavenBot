package dbhelp

import "fmt"

//ConditionGroup is group of SQL Condition, joined with logic operator
type ConditionGroup struct {
	ConLeft  *Condition
	LogOp    SQLLogOperator
	ConRight *Condition
}

//Condition is a SQL Condition
type Condition struct {
	Column *SQLColumn
	RelOp  SQLRelOperator
	Value  interface{}
}

//ToString parse to SQL transact the condition
func (c *Condition) ToString() string {
	pattern := "[%s] %s %v"
	if c.Column.Type == SQLTypeText {
		pattern = "[%s] %s '%v'"
	}

	if c.RelOp == Conts || c.RelOp == StartW || c.RelOp == EndW {
		value := fmt.Sprintf(string(c.RelOp), c.Value)
		return fmt.Sprintf("[%s] %s", c.Column.Name, value)
	}

	return fmt.Sprintf(pattern, c.Column.Name, c.RelOp, c.Value)
}

//ToString parse to SQL transact the group of conditions
func (g ConditionGroup) ToString() string {
	if g.ConRight == nil {
		return fmt.Sprintf("%s", g.ConLeft.ToString())
	}

	return fmt.Sprintf("%s %s %s", g.ConLeft.ToString(), g.LogOp, g.ConRight.ToString())
}
