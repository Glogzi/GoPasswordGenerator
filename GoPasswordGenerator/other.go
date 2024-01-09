package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func strToArray(strToConvert string) []string {
	var OutputArray []string
	for _, v := range strToConvert {
		OutputArray = append(OutputArray, string(v))
	}
	return OutputArray
}

func random(array []string) string {
	clear()
	print("how long password should be?\n>")
	buff := bufio.NewScanner(os.Stdin)
	var length int
	for {
		_, err := fmt.Scanln(&length)
		if err != nil {
			print("write a number...\n>")
			buff.Scan()
			continue
		}
		break
	}
	var arrayToOutput []string
	for i := 0; i < length; i++ {
		for {
			rngNum := rand.Intn(len(array))
			arrayToOutput = append(arrayToOutput, array[rngNum])
			break
		}
	}
	outputString := strings.Join(arrayToOutput, "")
	return outputString
}

func clear() {
	switch runtime.GOOS { //runtime.GOOS returns OS name
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
