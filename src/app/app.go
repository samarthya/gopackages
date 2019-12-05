package main

import (
	"fmt"
	"lib"
)

func init() {
	fmt.Println(" Main package: Main function (Init)")
}

/**
 * https://golang.org/doc/code.html
 * how to organize the code.
 */
func main() {
	fmt.Printf(" %s to Tutorial 1. \n Creating packages. \n %s\n", Message, lib.Greet)
}
