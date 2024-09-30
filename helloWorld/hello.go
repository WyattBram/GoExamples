package main

const (
	Spanish = "Spanish"
	French  = "French"

	Englishgreeting = "Hello, "
	Spanishgreeting = "Elode, "
	Frenchgreeting  = "Bonjure, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := Englishgreeting

	switch language {

	case Spanish:
		prefix = Spanishgreeting

	case French:
		prefix = Frenchgreeting

	}

	return prefix + name
}
