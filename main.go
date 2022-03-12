package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// Change this to your input pin, this example uses pin 4
const pinNum uint8 = 4

var pin = rpio.Pin(pinNum)

type Pulse struct {
	Value  uint8
	Length int64
}

func getData() []byte {
	var count1 int = 0
	var command []Pulse = make([]Pulse, 0)
	var binary = make([]byte, 0)
	binary = append(binary, 1)
	var previousValue uint8 = 0
	var value = uint8(pin.Read())

	for value == 1 {
		value = uint8(pin.Read())
		time.Sleep(time.Microsecond * 100)
	}
	startTime := time.Now()
	for {
		time.Sleep(time.Nanosecond * 50)
		if value != previousValue {
			now := time.Now()
			pulseLength := now.Sub(startTime)
			startTime = now
			command = append(command, Pulse{Value: previousValue, Length: pulseLength.Microseconds()})
		}
		if value == 1 {
			count1 += 1
		} else {
			count1 = 0
		}
		if count1 > 10000 {
			break
		}
		previousValue = value
		value = uint8(pin.Read())
	}
	for _, item := range command {
		if item.Value == 1 {
			if item.Length > 1000 {
				binary = append(binary, 0)
				binary = append(binary, 1)
			} else {
				binary = append(binary, 0)
			}
		}
	}
	if len(binary) > 34 {
		binary = binary[:34]
	}
	return binary
}

func main() {
	if err := rpio.Open(); err != nil {
		panic(err.Error())
	}
	pin.Input()
	fmt.Println("Please enter the name of the button you want to scan.")
	var buttonName string
	fmt.Scanln(&buttonName)
	fmt.Print("\x1b[1A")
	fmt.Printf("Press \x1b[1;32m%s\x1b[1;0m when commanded in order to scan.\n", buttonName)
	var resultMap map[string]uint8 = make(map[string]uint8)
	for i := 0; i < 10; i++ {
		fmt.Printf("Press \x1b[1;32m%s\x1b[1;0m now.\n", buttonName)
		data := getData()
		res := ""
		for _, v := range data {
			if v == 1 {
				res += "1"
			} else {
				res += "0"
			}
		}
		result := new(big.Int)
		result.SetString(res, 2)
		fmt.Printf("\x1b[1A[%-2d / 10] Result of \x1b[1;32m%s\x1b[1;0m: %x\n", i, buttonName, result)
		resultMap[fmt.Sprintf("%x", result)] += 1
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
