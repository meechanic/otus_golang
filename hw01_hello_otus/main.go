package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	const greeting string = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(greeting))
}
