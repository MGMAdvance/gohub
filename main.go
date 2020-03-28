package main

import (
	"fmt"

	"./user"
)

func main() {
	git := user.Github{}

	fmt.Print("Enter a Github username: ")

	var input string

	fmt.Scanln(&input)

	git.Search(input)

	fmt.Println(git.GetID())
}
