package db

import (
	"github.com/tecnologer/HellOrHeavenBot/db/dbhelp"
	hpr "github.com/tecnologer/HellOrHeavenBot/db/dbhelp"
	"github.com/tecnologer/HellOrHeavenBot/model"
)

var tableCResponseIsCreated bool

var customResTable = &hpr.SQLTable{
	Name: "CustomResponse",
	Columns: []*hpr.SQLColumn{
		dbhelp.NewPKCol("Id"),
		hpr.NewTextCol("Regex"),
		hpr.NewTextCol("Response"),
		hpr.NewIntCol("ResponseType"),
		hpr.NewTextNilCol("ChatID"),
		hpr.NewIntNilCol("Author"),
	},
}

//RetrieveCustomResponses returns an array with the CustomReponse for the selected chat
func RetrieveCustomResponses(chatID string) ([]*model.CustomResponse, error) {
	err := customResTable.Create()
	if err != nil {
		return nil, err
	}

	conditions := []*hpr.ConditionGroup{
		&hpr.ConditionGroup{
			ConLeft: &hpr.Condition{
				Column: customResTable.GetColByName("ChatID"),
				RelOp:  hpr.Eq,
				Value:  chatID,
			},
			LogOp: hpr.Or,
			ConRight: &hpr.Condition{
				Column: customResTable.GetColByName("ChatID"),
				RelOp:  hpr.Eq,
				Value:  nil,
			},
		},
	}

	rows, err := customResTable.ExecSelectCols([]string{"Regex", "Response", "ResponseType"}, conditions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	responses := make([]*model.CustomResponse, 0)
	for rows.Next() {
		res := new(model.CustomResponse)

		err = rows.Scan(&res.Regex, &res.Response, &res.ResponseType)
		if err != nil {
			return nil, err
		}

		responses = append(responses, res)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return responses, nil
}

func InsertCustomResponse(cRes *model.CustomResponse) error {
	err := customResTable.Create()
	if err != nil {
		return err
	}

	var chatID interface{}

	if cRes.ChatID != "" && cRes.ChatID != "0" {
		chatID = cRes.ChatID
	}

	return customResTable.Insert(cRes.Regex, cRes.Response, cRes.ResponseType, chatID, cRes.Author)
}
