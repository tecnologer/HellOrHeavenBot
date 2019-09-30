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
		hpr.NewIntNilCol("ChatID"),
		hpr.NewIntNilCol("Author"),
	},
}

func init() {
	customResTable.Create()
}

const (
	queryGetCResponseCommand = "SELECT [CommandID], [Response],[Type] FROM [%s] WHERE [CommandID] = %d AND [Language] = '%s';"
	queryCInsertResponse     = "INSERT INTO [%s] (Type, CommandID, Response, Language) VALUES (%d, %d, '%s', '%s', %t);"
)

//RetrieveCustomResponses returns an array with the CustomReponse for the selected chat
func RetrieveCustomResponses(chatID int) ([]*model.CustomResponse, error) {

	return nil, nil
}
