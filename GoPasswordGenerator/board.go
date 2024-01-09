package main

import (
	"fmt"
	"reflect"
	"strings"
)

func board() []string {
	//colors
	const red = "\033[1;31m"
	const reset = "\033[1;39m"
	const green = "\033[1;32m"

	//colored X and V
	const redX = red + "X" + reset
	const greenV = green + "V" + reset

	//current character set mode
	uclIsSelected := redX
	lclIsSelected := redX
	ssIsSelected := redX
	nbIsSelected := redX

	//character sets
	const uppercaseLetters = "QWERTYUIOPASDFGHJKLZXCVBNM"
	const lowercaseLetters = "qwertyuiopasdfghjklzxcvbnm"
	const specialSigns = "!@#$%^&*()-=+_[];',./<>?:{}\\|\""
	const numbers = "1234567890"

	//converting character sets to an array
	uppercaseLettersArray := strToArray(uppercaseLetters)
	lowercaseLettersArray := strToArray(lowercaseLetters)
	specialSignsArray := strToArray(specialSigns)
	numbersArray := strToArray(numbers)

	var everythingArray []string

	message := ""
	//displaying the board
	for {
		fmt.Println(message)
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
		usrInput = strings.ToUpper(usrInput)

		//changing current character set mode
		switch usrInput {
		case "UCL":
			if uclIsSelected == greenV {
				message = red + "already selected" + reset
				clear()
				continue
			}
			for _, v := range uppercaseLettersArray {
				everythingArray = append(everythingArray, v)
				uclIsSelected = greenV
			}
		case "LCL":
			if lclIsSelected == greenV {
				message = red + "already selected" + reset
				clear()
				continue
			}
			for _, v := range lowercaseLettersArray {
				everythingArray = append(everythingArray, v)
				lclIsSelected = greenV
			}
		case "SS":
			if ssIsSelected == greenV {
				message = red + "already selected" + reset
				clear()
				continue
			}
			for _, v := range specialSignsArray {
				everythingArray = append(everythingArray, v)
				ssIsSelected = greenV
			}
		case "NB":
			if nbIsSelected == greenV {
				message = red + "already selected" + reset
				clear()
				continue
			}
			for _, v := range numbersArray {
				everythingArray = append(everythingArray, v)
				nbIsSelected = greenV
			}
		case "RESET":
			everythingArray = everythingArray[:0] //resets array
			uclIsSelected = redX
			lclIsSelected = redX
			ssIsSelected = redX
			nbIsSelected = redX
		case "CONFIRM":
			//array to compare it to everythingArray to handle error when everythingArray is empty
			var emptyArray []string
			if !reflect.DeepEqual(everythingArray, emptyArray) {
				return everythingArray
			}
			message = red + "I need something to generate password..." + reset
		default:
			message = red + "wrong input" + reset
		}
		//the end of switch
		clear()
	}
}
