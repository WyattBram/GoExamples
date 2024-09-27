package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris"))
}

const greeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		return greeting + "World"
	}
	return greeting + name
}
