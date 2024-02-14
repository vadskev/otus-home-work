package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	line := "Hello, OTUS!"
	lineReverse := reverse.String(line)
	fmt.Println(lineReverse)
}
