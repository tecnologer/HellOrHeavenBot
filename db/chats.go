package db

import (
	hpr "github.com/tecnologer/HellOrHeavenBot/db/dbhelp"
	m "github.com/tecnologer/HellOrHeavenBot/model"
)

var tableChatsIsCreated bool
var chatsTable = &hpr.SQLTable{
	Name: "ChatLog",
	Columns: []*hpr.SQLColumn{
		hpr.NewPKCol("Id"),
		hpr.NewIntCol("ChatID"),
		hpr.NewTextCol("Name"),
	},
}

const (
// querySearchChatByName = "SELECT [ChatId], [Name] FROM [%s] WHERE [Name] like '%%%s%%'"
//querySearchChatByID = "SELECT [ChatId], [Name] FROM [%s] WHERE [ChatId] = %d"
// queryGetChatByName = "SELECT [ChatId], [Name] FROM [%s] WHERE [Name] = '%s'"
// queryUpdateChat    = "UPDATE [%s] SET [Name]=%s WHERE [ChatId]=%d;"
// queryInsertChat    = "INSERT INTO [%s] (ChatId, Name) VALUES (%d, '%s');"
)

func init() {
	chatsTable.Create()
}

//InsertOrUpdateChat Inserts a new chat or update it if is exists
func InsertOrUpdateChat(chat *m.Chat) error {
	err := chatsTable.Create()
	if err != nil {
		return err
	}

	conditions := []*hpr.ConditionGroup{
		&hpr.ConditionGroup{
			ConLeft: &hpr.Condition{
				Column: chatsTable.Columns[1],
				RelOp:  hpr.Eq,
				Value:  chat.ID,
			},
		},
	}

	rows, err := chatsTable.ExecSelectCols([]string{"ChatId", "Name"}, conditions)

	if err != nil {
		return err
	}

	defer rows.Close()
	var registeredChat *m.Chat

	for rows.Next() {
		registeredChat = new(m.Chat)
		err = rows.Scan(&registeredChat.ID, &registeredChat.Name)
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	var isInsert = registeredChat == nil

	if isInsert {
		err = chatsTable.Insert(chat.ID, chat.Name)
	} else {

		if chat.Name == registeredChat.Name {
			return nil
		}

		updateValues := map[string]interface{}{
			"Name":   chat.Name,
			"ChatId": chat.ID,
		}

		err = chatsTable.Update(updateValues, conditions)
	}

	if err != nil {
		return err
	}
	return nil
}
