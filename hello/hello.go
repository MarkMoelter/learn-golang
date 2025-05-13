package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case "":
		prefix = englishHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}
