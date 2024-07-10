package piscine

import "os"

func PrintError(err error) {
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