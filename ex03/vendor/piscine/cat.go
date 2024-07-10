package piscine

import "os"

func PrintError(err error) {
	os.Stdout.Write([]byte("ERROR: "))
	os.Stdout.Write([]byte(err.Error()))
	os.Stdout.Write([]byte("\n"))
}

func DisplayFile(file *os.File) {
	const bufSize = 1024
	buf := make([]byte, bufSize)
	for {
		n, err := file.Read(buf)
		if err == nil && n != 0 {
			os.Stdout.Write(buf[:n])
		} else if err.Error() == "EOF" || n == 0 {
			break
		} else if err != nil {
			PrintError(err)
			break
		}
	}
}

func Cat() {
	len := 0
	for range os.Args {
		len++
	}
	if len < 2 {
		DisplayFile(os.Stdin)
	} else if len == 2 {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			PrintError(err)
			os.Exit(1)
		} else {
			defer file.Close()
			DisplayFile(file)
		}
	} else {
		for i := 1; i < len; i++ {
			filename := os.Args[i]
			file, err := os.Open(filename)
			if err != nil {
				PrintError(err)
			} else {
				defer file.Close()
				DisplayFile(file)
			}
		}
	}
}
