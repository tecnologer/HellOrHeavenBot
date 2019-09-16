package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/tecnologer/HellOrHeavenBot/model"
	m "github.com/tecnologer/HellOrHeavenBot/model"
)

var tableChatsIsCreated bool

const (
	tableNameChats        = "ChatLog"
	queryCreateTableChats = `CREATE TABLE IF NOT EXISTS [%s] (
		[Id] integer not null primary key AUTOINCREMENT,
		[ChatId] integer not null,
		[Name] text not null
	)`
	querySearchChatByName = "SELECT [ChatId], [Name] FROM [%s] WHERE [Name] like '%%%s%%'"
	querySearchChatByID   = "SELECT [ChatId], [Name] FROM [%s] WHERE [ChatId] = %d"
	queryGetChatByName    = "SELECT [ChatId], [Name] FROM [%s] WHERE [Name] = '%s'"
	queryUpdateChat       = "UPDATE [%s] SET [Name]=%s WHERE [ChatId]=%d;"
	queryInsertChat       = "INSERT INTO [%s] (ChatId, Name) VALUES (%d, '%s');"
)

func init() {
	createTableChats()
}

func createTableChats() {
	log.Printf("creating table %s\n", tableNameChats)
	if tableChatsIsCreated {
		return
	}

	err := execQueryNoResult(queryf(queryCreateTableChats, tableNameChats))
	tableChatsIsCreated = err == nil
	if !tableChatsIsCreated {
		log.Println(err)
	} else {
		log.Printf("table %s is created\n", tableNameChats)
	}

}

//InsertOrUpdateChat Inserts a new chat or update it if is exists
func InsertOrUpdateChat(chat *model.Chat) error {
	createTableStats()

	getChatByID := queryf(querySearchChatByID, tableNameChats, chat.ID)

	rows, err := execQuery(getChatByID)

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
	var tmpQuery query

	if isInsert {

		tmpQuery = queryf(queryInsertChat, tableNameChats, chat.ID, chat.Name)
	} else {

		if chat.Name == registeredChat.Name {
			return nil
		}

		tmpQuery = queryf(queryUpdateChat, tableNameChats, chat.Name, chat.ID)
	}

	err = execQueryNoResult(tmpQuery)
	if err != nil {
		return err
	}
	return nil
}
