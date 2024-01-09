package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	switch runtime.GOOS { //runtime.GOOS returns OS
	case "windows":
		command := exec.Command("cmd", "/c", "cls")
		command.Stdout = os.Stdout
		command.Run()
	case "linux":
		command := exec.Command("clear")
		command.Stdout = os.Stdout
		command.Run()
	}

}
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
	red := "\033[1;31m"
	reset := "\033[1;39m"
	green := "\033[1;32m"
	redX := red + "X" + reset
	greenV := green + "V" + reset
	uclIsSelected := redX
	lclIsSelected := redX
	ssIsSelected := redX
	nbIsSelected := redX
	uppercaseLetters := "QWERTYUIOPASDFGHJKLZXCVBNM"
	lowercaseLetters := "qwertyuiopasdfghjklzxcvbnm"
	specialSigns := "!@#$%^&*()-=[];',./<>?:{}\\|\""
	numbers := "1234567890"
	uppercaseLettersArray := strToArray(uppercaseLetters)
	lowercaseLettersArray := strToArray(lowercaseLetters)
	specialSignsArray := strToArray(specialSigns)
	numbersArray := strToArray(numbers)
	var everythingArray []string
	message := ""
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
			for _, v := range numbersArray {
				everythingArray = append(everythingArray, v)
				ssIsSelected = greenV
			}
		case "NB":
			if nbIsSelected == greenV {
				message = red + "already selected" + reset
				clear()
				continue
			}
			for _, v := range specialSignsArray {
				everythingArray = append(everythingArray, v)
				nbIsSelected = greenV
			}
		case "reset":
			everythingArray = everythingArray[:0] //resets array
			uclIsSelected = redX
			lclIsSelected = redX
			ssIsSelected = redX
			nbIsSelected = redX
		case "next":
			return everythingArray
		default:
			message = red + "wrong input" + reset
		}
		//the end of switch
		clear()
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
