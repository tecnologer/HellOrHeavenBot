package model

//ResponseType enum for type of the responses
type ResponseType byte

//Response is the struct to manage the responses for the commands
type Response struct {
	Content   string       `json:"a"`
	Type      ResponseType `json:"at"`
	CommandID byte         `json:"t"`
	Language  string
}

const (
	//Text is for response as plain text
	Text ResponseType = iota
	//Sticker is the id of a sticker
	Sticker
	//Gif is the id of a gif
	Gif
	//Photo is a document of image
	Photo
	//Broadcast tbd
	Broadcast
)
