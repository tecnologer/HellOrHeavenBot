package db

import (
	"fmt"

	"github.com/tecnologer/HellOrHeavenBot/resources"

	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/model"
	m "github.com/tecnologer/HellOrHeavenBot/model"
)

var tableResponsesIsCreated bool

const (
	tableNameResponses        = "Response"
	queryCreateTableResponses = `CREATE TABLE IF NOT EXISTS [%s] (
		[Id] integer not null primary key AUTOINCREMENT,
		[Type] integer not null,
		[CommandID] integer not null,
		[Response] text not null
	);`
	queryGetResponseCommand = "SELECT [CommandID], [Response],[Type] FROM [%s] WHERE [CommandID] = %d;"
	queryInsertResponse     = "INSERT INTO [%s] (Type, CommandID, Response) VALUES (%d, %d, '%s');"
)

func init() {
	createTableResponses()
}

func createTableResponses() {
	log.Printf("creating table %s\n", tableNameResponses)
	if tableResponsesIsCreated {
		return
	}

	err := execQueryNoResult(queryf(queryCreateTableResponses, tableNameResponses))
	tableResponsesIsCreated = err == nil
	if !tableResponsesIsCreated {
		log.Println(err)
	} else {
		log.Printf("table %s is created\n", tableNameResponses)
	}

}

//InsertResponse creates new record of response in Response
func InsertResponse(res *model.Response) error {
	createTableResponses()

	tmpQuery := queryf(queryInsertResponse, tableNameResponses, res.Type, res.CommandID, res.Content)

	err := execQueryNoResult(tmpQuery)
	if err != nil {
		return err
	}
	return nil
}

//GetResponseByCommand select a response (random) assigned to the command
func GetResponseByCommand(comID int) (*model.Response, error) {
	tmpQuery := queryf(queryGetResponseCommand, tableNameResponses, comID)
	rows, err := execQuery(tmpQuery)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var responses []*model.Response
	for rows.Next() {
		response := new(m.Response)

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

	selection := resources.GetRandomIntFromRange(0, len(responses)-1)

	return responses[selection], nil
}