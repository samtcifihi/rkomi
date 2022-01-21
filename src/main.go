package main

import (
	_ "flag"
	"fmt"
	"os"
)

func main() {
	test := os.Args[0]
	fmt.Println(test, "ran successfuly.") // DEBUG
}
