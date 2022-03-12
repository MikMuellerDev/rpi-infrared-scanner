package main

import (
	"fmt"

	"github.com/MikMuellerDev/rpiif"
)

const pin = 4

func main() {
	scanner := rpiif.Scanner
	if err := scanner.Setup(pin); err != nil {
		panic(err.Error())
	}
	fmt.Println("Please enter the name of the button you want to scan.")
	var buttonName string
	fmt.Scanln(&buttonName)
	fmt.Print("\x1b[1A")
	fmt.Printf("Press \x1b[1;32m%s\x1b[1;0m when commanded in order to scan.\n", buttonName)
	var resultMap map[string]uint8 = make(map[string]uint8)
	for i := 0; i < 10; i++ {
		fmt.Printf("Press \x1b[1;32m%s\x1b[1;0m now.\n", buttonName)
		result, err := scanner.Scan()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("\x1b[1A[%-2d / 10] Result of \x1b[1;32m%s\x1b[1;0m: %s\n", i, buttonName, result)
		resultMap[result] += 1
	}
	var maxOccurences uint8 = 0
	for _, occurrence := range resultMap {
		if occurrence > maxOccurences {
			maxOccurences = occurrence
		}
	}
	for value, occurrence := range resultMap {
		if occurrence == maxOccurences {
			fmt.Printf("\x1b[1AThe command for \x1b[1;32m%s\x1b[1;0m is likely: \x1b[1;33m%s\x1b[1;0m			\n", buttonName, value)
			break
		}
	}
}
