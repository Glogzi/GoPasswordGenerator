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
	/*lowercaseLetters := "qwertyuiopasdfghjklzxcvbnm"
	specialSigns := "!@#$%^&*()-=[];',./<>?:{}\\|\""
	numbers := "1234567890"*/
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
			strToArray(uppercaseLetters)
		}

	}

}
func main() {
	fmt.Println("welcome to password generator, please \nselect from what password will be generated")
}
