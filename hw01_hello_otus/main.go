package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	const greeting string = "Hello, OTUS!"
	fmt.Println(reverse.String(greeting))
}
