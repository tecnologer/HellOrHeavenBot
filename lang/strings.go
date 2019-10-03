package lang

//messages has the strings for static messages, groupped by language
var messages = languageList{
	"en": map[string]string{
		"ticketsNameRequired":        "The user name is required",
		"genericFail":                "Something was wrong!",
		"requestResponseContent":     "The next message you sent will be taken as response. You can use text, sticker or gif",
		"responseStored":             "Response stored correctly",
		"customResponseStored":       "Done!",
		"genericCancel":              "Action canceled",
		"cResponseRegexReq":          "The regex is required. `/customanswer <regex>`",
		"customResponseStoredFailed": "Something went wrong. Try again",
	},
	"es": map[string]string{
		"ticketsNameRequired":        "El nombre del condenado es requerido",
		"genericFail":                "falio ferga!",
		"requestResponseContent":     "El siguiente mensaje que envies se tomara como respuesta. Puedes usar texto, sticker o gif",
		"responseStored":             "Listoooo!!",
		"customResponseStored":       "Listoooo!!",
		"genericCancel":              "Accion cancelada",
		"cResponseRegexReq":          "La expresion regular es requerida. `/customanswer <regex>`",
		"customResponseStoredFailed": "Algo valio camote, intenta de nuevo",
	},
}
