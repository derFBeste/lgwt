package hello

import "fmt"

// type Language struct {
// 	English string
// }

// var language = &Hello.Params{
// 	English: "English",
// }

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// func HelloEnglish(name string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	return englishHelloPrefix + name
// }

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix

	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
