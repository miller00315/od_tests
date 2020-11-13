package main

import (
	"fmt"
	"tests-introduction/addresses"
)

func main() {
	typeAddress := addresses.TypeAddress("Rodovia dos imigrantes")
	fmt.Println(typeAddress)
}
