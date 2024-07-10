package piscine

import "os"

const FORE = 0
const BACK = 1
const TENN = 2

func PrintError(err error) {
	os.Stderr.Write([]byte("ERROR: "))
	os.Stderr.Write([]byte(err.Error()))
	os.Stderr.Write([]byte("\n"))
}

func Get10LineN(str string) uint64 {
	len := 0
	for range str {
		len++
	}
	count := 0
	for i := len - 1; i >= 0; i-- {
		if str[i] == '\n' {
			count++
		}
		if count == 11 {
			return uint64(i+1)
		}
	}
	return 0
}

func convertNumber(str string) (uint64, bool) {
	var n uint64
	n = 0
	max := ^uint64(0)
	for _, c := range str {
		if (c < '0' || c > '9') || (max-uint64(c-'0')/10 < n) {
			return n, false
		}
		n = n*10 + uint64(c-'0')
	}
	return n, true
}

func GetDataFiles(len int, args []string, flag int, n uint64) (string, bool) {
	res := ""
	success := true
	count := 0
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			res += "open "
			res += arg
			res += ": no such file or directory\n"
			success = false
		} else {
			count++
			if len > 1 && count != 1 {
				res += "\n"
			}
			if len > 1 {
				res += "==> " + arg + " <==\n"
			}
			defer file.Close()
			data, err := os.ReadFile(arg)
			if err != nil && err.Error() != "EOF" {
				res += err.Error()
				success = false
			} else {
				var data_len uint64
				data_len = 0
				for range data {
					data_len++
				}
				if flag == TENN {
					n = Get10LineN(string(data))
				}
				if flag == FORE {
					n = n - 1
				}
				if flag == BACK {
					n = data_len - n
				}
				if data_len < n {
					n = data_len
				}
				if n < 0 {
					n = 0
				}
				res += string(data)[n:]
			}
		}
	}
	return res, success
}

func GetDataStdin(len int, flag int, n uint64) (string, bool) {
	res := ""
	data := ""
	success := true
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil && err.Error() != "EOF" {
			res += err.Error()
			success = false
			break
		} else if n == 0 {
			break
		} else {
			data += string(buf[:n])
		}
	}
	if success == false {
		return res, success
	}
	var data_len uint64
	data_len = 0
	for range data {
		data_len++
	}
	if flag == TENN {
		n = Get10LineN(string(data))
	}
	if flag == FORE {
		n = n - 1
	}
	if flag == BACK {
		n = data_len - n
	}
	if data_len < n {
		n = data_len
	}
	if n < 0 {
		n = 0
	}
	res += string(data)[n:]
	return res, success
}

func handleOption() (int, int, uint64, []string) {
	len := 0
	flag := TENN
	var n uint64
	n = 0
	nflag := true
	args := os.Args[1:]
	for range args {
		len++
	}
	if len > 0 {
		if args[0] == "-c" {
			args = args[1:]
			len--
			if len > 0 {
				num_str := args[0]
				flag = BACK
				if num_str[0] == '+' {
					flag = FORE
					num_str = num_str[1:]
				} else if num_str[0] == '-' {
					flag = BACK
					num_str = num_str[1:]
				}
				n, nflag = convertNumber(num_str)
				if !nflag {
					os.Stderr.Write([]byte("tail: invalid number of bytes: ‘"))
					os.Stderr.Write([]byte(num_str))
					os.Stderr.Write([]byte("’\n"))
					os.Exit(1)
				}
				args = args[1:]
				len--
			}
		} else {
			flag = TENN
		}
	}
	return len, flag, n, args
}

func Ztail() {
	len, flag, n, args := handleOption()
	success := true
	data := ""
	tmp_success := true
	if len == 0 {
		data, tmp_success = GetDataStdin(len, flag, n)
	} else {
		data, tmp_success = GetDataFiles(len, args, flag, n)
	}
	success = success && tmp_success
	os.Stdout.Write([]byte(data))
	if !success {
		os.Exit(1)
	}
}
