package main

import (
	"fmt"
	"os/exec"
)

func strToArray(strToConvert string) [28]string {
	var OutputArray [28]string
	for i, v := range strToConvert {
		if string(v) != "" {
			OutputArray[i] = string(v)
		}
	}
	return OutputArray
}
func board() []string {
	var uclIsSelected = "X"
	var lclIsSelected = "X"
	var ssIsSelected = "X"
	var nbIsSelected = "X"
	uppercaseLetters := "QWERTYUIOPASDFGHJKLZXCVBNM"
	lowercaseLetters := "qwertyuiopasdfghjklzxcvbnm"
	specialSigns := "!@#$%^&*()-=[];',./<>?:{}\\|\""
	numbers := "1234567890"
	uppercaseLettersArray := strToArray(uppercaseLetters)
	lowercaseLettersArray := strToArray(lowercaseLetters)
	specialSignsArray := strToArray(specialSigns)
	numbersArray := strToArray(numbers)
	var everythingArray []string
	for {
		fmt.Printf("%v UCL - uppercase letters\n"+
			"%v LCL - lower case letters\n"+
			"%v SS - special signs\n"+
			"%v NB - numbers", uclIsSelected, lclIsSelected, ssIsSelected, nbIsSelected)
		print("\n>")
		var usrInput string
		_, err := fmt.Scanln(&usrInput)
		if err != nil {
			continue
		}
		switch usrInput {
		case "UCL":
			if uclIsSelected == "V" {
				exec.Command("cls || clear")
				println("already selected")
				continue
			}
			for _, v := range uppercaseLettersArray {
				everythingArray = append(everythingArray, v)
				uclIsSelected = "V"
			}
		case "LCL":
			if lclIsSelected == "V" {
				exec.Command("cls || clear")
				println("already selected")
				continue
			}
			for _, v := range lowercaseLettersArray {
				everythingArray = append(everythingArray, v)
				lclIsSelected = "V"
			}
		case "SS":
			if ssIsSelected == "V" {
				exec.Command("cls || clear")
				println("already selected")
				continue
			}
			for _, v := range numbersArray {
				everythingArray = append(everythingArray, v)
				ssIsSelected = "V"
			}
		case "NB":
			if nbIsSelected == "V" {
				exec.Command("cls || clear")
				println("already selected")
				continue
			}
			for _, v := range specialSignsArray {
				everythingArray = append(everythingArray, v)
				nbIsSelected = "V"
			}
		case "reset":
			everythingArray = everythingArray[:0]
			uclIsSelected = "X"
			lclIsSelected = "X"
			ssIsSelected = "X"
			nbIsSelected = "X"
		case "next":
			return everythingArray
		default:
			println("wrong input!!!")
		}
		//the end of switch
		exec.Command("cls || clear")
	}

}
func main() {
	fmt.Println("welcome to password generator, please \nselect from what password will be generated")
	fmt.Println("if added something you don't want, you can reset by typing \"reset\"")
	fmt.Println("if you want to go to next step type \"confirm\"")
	x := board()
	for _, v := range x {
		fmt.Printf("%v, ", v)
	}
}
