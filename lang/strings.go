package lang

//messages has the strings for static messages, groupped by language
var messages = languageList{
	"en": map[string]string{
		"ticketsNameRequired":    "The user name is required",
		"genericFail":            "Something was wrong!",
		"requestResponseContent": "The next message you sent will be taken as response. You can use text, sticker or gif",
	},
	"es": map[string]string{
		"ticketsNameRequired":    "El nombre del condenado es requerido",
		"genericFail":            "falio ferga!",
		"requestResponseContent": "El siguiente mensaje que envies se tomara como respuesta. Puedes user texto, sticker o gif",
	},
}
