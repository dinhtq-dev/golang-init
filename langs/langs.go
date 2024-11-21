package langs

var Messages map[string]string

var CurrentLanguage string

func InitLanguage() {
	switch CurrentLanguage {
	case "en":
		Messages = enMessages
	default:
		Messages = enMessages
	}
}
