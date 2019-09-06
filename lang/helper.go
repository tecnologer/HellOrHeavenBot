package lang

const defaultLan = "en"

//languageList group of messages by language code
type languageList map[string]map[string]string

//GetMessagesByLanguage returns the list of messages according the language code
func GetMessagesByLanguage(lang string) map[string]string {
	if lang, exists := messages[lang]; exists {
		return lang
	}

	return messages[defaultLan]
}
