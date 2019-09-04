package model

import "regexp"

//CustomResponse is a struct to store the rules based on Regex
type CustomResponse struct {
	Regex        *regexp.Regexp `json:"regex"`
	Response     string         `json:"answer"`
	ResponseType ResponseType   `json:"type"`
	ChatID       int            `json:"chat_id"`
	Author       int            `json:"author"`
}
