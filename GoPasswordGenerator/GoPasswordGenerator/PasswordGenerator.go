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
	var ssIsSelected = "x"
	var nbIsSelected = "X"
	uppercaseLetters := "QWERTYUIOPASDFGHJKLZXCVBNM"
	lowercaseLetters := "qwertyuiopasdfghjklzxcvbnm"
	/*specialSigns := "!@#$%^&*()-=[];',./<>?:{}\\|\""
	numbers := "1234567890"*/
	uppercaseLettersArray := strToArray(uppercaseLetters)
	lowercaseLettersArray := strToArray(lowercaseLetters)
	/*specialSignsArray := strToArray(specialSigns)
	numbersArray := strToArray(numbers)*/
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
			for _, v := range uppercaseLettersArray {
				everythingArray = append(everythingArray, v)
			}
		case "LCL":
			for _, v := range lowercaseLettersArray {
				everythingArray = append(everythingArray, v)
			}
		}

	}

}
func main() {
	fmt.Println("welcome to password generator, please \nselect from what password will be generated")
	fmt.Println("if added something you didn't want, you can reset by typing \"reset\"")
}
