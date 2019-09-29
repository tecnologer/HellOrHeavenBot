package db

import (
	"fmt"

	hpr "github.com/tecnologer/HellOrHeavenBot/db/dbhelp"
	"github.com/tecnologer/HellOrHeavenBot/model"
	"github.com/tecnologer/HellOrHeavenBot/resources"
)

var responsesTable = &hpr.SQLTable{
	Name: "Response",
	Columns: []*hpr.SQLColumn{
		hpr.NewPKCol("Id"),
		hpr.NewIntCol("Type"),
		hpr.NewIntCol("CommandID"),
		hpr.NewTextCol("Response"),
		hpr.NewTextCol("Language"),
		hpr.NewIntNilCol("IsAnimated"),
	},
}

//InsertResponse creates new record of response in Response
func InsertResponse(res *model.Response) error {
	err := responsesTable.Create()
	if err != nil {
		return err
	}

	err = responsesTable.Insert(res.Type, res.CommandID, res.Content, res.Language, res.IsAnimated)
	if err != nil {
		return err
	}
	return nil
}

//GetResponseByCommand select a response (random) assigned to the command
func GetResponseByCommand(comID int, lang string) (*model.Response, error) {
	err := responsesTable.Create()
	if err != nil {
		return nil, err
	}

	conditions := []*hpr.ConditionGroup{
		&hpr.ConditionGroup{
			ConLeft: &hpr.Condition{
				Column: chatsTable.GetColByName("CommandID"),
				RelOp:  hpr.Eq,
				Value:  comID,
			},
			LogOp: hpr.And,
			ConRight: &hpr.Condition{
				Column: chatsTable.GetColByName("Language"),
				RelOp:  hpr.Eq,
				Value:  lang,
			},
		},
	}

	rows, err := responsesTable.ExecSelectCols([]string{}, conditions)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var responses []*model.Response
	for rows.Next() {
		response := new(model.Response)

		err = rows.Scan(&response.CommandID, &response.Content, &response.Type)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(responses) == 0 {
		return nil, fmt.Errorf("no responses for the selected command")
	}

	if len(responses) == 1 {
		return responses[0], nil
	}

	selection := resources.GetRandomIntFromRange(0, len(responses)-1)

	return responses[selection], nil
}
