package main

import (
	"os"
	"piscine"
)

func main() {
	len := 0
	for range os.Args {
		len++
	}
	if len < 2 {
		os.Stdout.Write([]byte("File name missing\n"))
	} else if len > 2 {
		os.Stdout.Write([]byte("Too many arguments\n"))
	} else {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			piscine.PrintError(err)
		}
		defer file.Close()
		piscine.DisplayFile(file)
	}
}
