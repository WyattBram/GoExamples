package main

import "fmt"
import "strings"

func testCompare() {
	other := "e"
	fmt.Println(strings.Compare("e", other))
	fmt.Println(strings.Compare("f", other))
	fmt.Println(strings.Compare("a", other))
	fmt.Println(strings.Compare("E", other))
}

func main() {
	testCompare()
}
