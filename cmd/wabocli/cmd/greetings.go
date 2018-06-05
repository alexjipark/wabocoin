package cmd

func Greeting (lang string) string {
	switch lang {
	case "pl":
		return "Czesc"
	case "es":
		return "Hola"
	default:
		return "Hello"
	}
}
