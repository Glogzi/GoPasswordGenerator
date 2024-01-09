package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to password generator, please \nselect from what password will be generated")
	fmt.Println("if added something you don't want, you can reset by typing \"reset\"")
	fmt.Println("if you want to go to next step type \"confirm\"")
	passwordSigns := board()

	password := random(passwordSigns)
	println("your password: ", password)
	fmt.Scanln()
}
