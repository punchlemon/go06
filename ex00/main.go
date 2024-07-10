package main

import (
	"os"
	"ft"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := 0
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	for range os.Args[1:] {
		lengthOfArg++
	}
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
