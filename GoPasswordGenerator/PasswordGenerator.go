package main

import (
	"fmt"
	"log"
)

func strToArray(strToConvert string) []string {
	var outputArray []string
	for i, v := range strToConvert {
		outputArray[i] = string(v)
	}
	return outputArray
}
func board() {
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
		fmt.Printf("%v UCL - uppercase letters\n", uclIsSelected+
			"%v LCL - lower case letters\n", lclIsSelected+
			"%v SS - special signs\n", ssIsSelected+
			"%v NB - numbers", nbIsSelected)
		var usrInput string
		_, err := fmt.Scan(&usrInput)
		if err != nil {
			log.Fatal(err)
		}
		switch usrInput {
		case "UCL":
			if uclIsSelected == "V" {
				println("already selected")
				continue
			}
			for _, v := range uppercaseLettersArray {
				everythingArray = append(everythingArray, v)
				uclIsSelected = "V"
			}
		case "LCL":
			if lclIsSelected == "V" {
				println("already selected")
				continue
			}
			for _, v := range lowercaseLettersArray {
				everythingArray = append(everythingArray, v)
				uclIsSelected = "V"
			}
		case "SS":
			if ssIsSelected == "V" {
				println("already selected")
				continue
			}
			for _, v := range numbersArray {
				everythingArray = append(everythingArray, v)
				uclIsSelected = "V"
			}
		case "NB":
			if nbIsSelected == "V" {
				println("already selected")
				continue
			}
			for _, v := range specialSignsArray {
				everythingArray = append(everythingArray, v)
				uclIsSelected = "V"
			}
		case "reset":
			everythingArray = everythingArray[:0]
			uclIsSelected = "X"
			lclIsSelected = "X"
			ssIsSelected = "X"
			nbIsSelected = "X"
		default:
			println("wrong input!!!")

		}

	}

}
func main() {
	fmt.Println("welcome to password generator, please \nselect from what password will be generated")
	fmt.Println("if added something you don't want, you can reset by typing \"reset\"")
	fmt.Println("if you want to go to next step type \"confirm\"")
}
