package test

import (
	"testing"

	hpr "github.com/tecnologer/HellOrHeavenBot/db/dbhelp"
)

type CondTest struct {
	input  hpr.Condition
	output string
}

func TestConditionToString(t *testing.T) {
	inputTest := []CondTest{
		CondTest{
			input: hpr.Condition{
				Column: hpr.NewIntCol("Id"),
				RelOp:  hpr.Eq,
				Value:  1,
			},
			output: "[Id] == 1",
		},
		CondTest{
			input: hpr.Condition{
				Column: hpr.NewTextCol("Id"),
				RelOp:  hpr.Eq,
				Value:  1,
			},
			output: "[Id] == '1'",
		},
		CondTest{
			input: hpr.Condition{
				Column: hpr.NewTextCol("Id"),
				RelOp:  hpr.Conts,
				Value:  1,
			},
			output: "[Id] like '%1%'",
		},
		CondTest{
			input: hpr.Condition{
				Column: hpr.NewTextCol("Id"),
				RelOp:  hpr.StartW,
				Value:  1,
			},
			output: "[Id] like '1%'",
		},
		CondTest{
			input: hpr.Condition{
				Column: hpr.NewTextCol("Id"),
				RelOp:  hpr.EndW,
				Value:  1,
			},
			output: "[Id] like '%1'",
		},
	}

	for _, it := range inputTest {
		str := it.input.ToString()

		if str != it.output {
			t.Log("Failed", str, "Expected", it.output)
			t.Fail()
		}

		t.Log(it.output, str)
	}
}
