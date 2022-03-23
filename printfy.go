package printfy

import (
	"fmt"
	"strconv"
)

type any = interface{}

func Printfy(stri string, args ...any) {
	arg := false
	str := ""

	for i := 0; i < len(stri); i++ {
		switch string(stri[i]) {
		case "{":
			arg = true
			continue
		case "}":
			arg = false
			continue
		}

		if !arg {
			str = str + string(stri[i])
		} else {
			args_index, _ := strconv.Atoi(string(stri[i]))
			str = str + fmt.Sprintf("%v", args[args_index])
			args_index++
		}
	}
	fmt.Print(str)
}

func Stringfy(stri string, args ...any) string {
	arg := false
	str := ""

	for i := 0; i < len(stri); i++ {
		switch string(stri[i]) {
		case "{":
			arg = true
			continue
		case "}":
			arg = false
			continue
		}

		if !arg {
			str = str + string(stri[i])
		} else {
			args_index, _ := strconv.Atoi(string(stri[i]))
			str = str + fmt.Sprintf("%v", args[args_index])
			args_index++
		}
	}
	return str
}
