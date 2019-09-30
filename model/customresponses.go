package model

import "encoding/json"

//CustomResponse is a struct to store the rules based on Regex
type CustomResponse struct {
	Regex        string       `json:"regex"`
	Response     string       `json:"answer"`
	ResponseType ResponseType `json:"type"`
	ChatID       json.Number  `json:"chat_id,Number"`
	Author       int          `json:"author"`
}
