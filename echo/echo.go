package echo

import (
	"strconv"
	"strings"
)

func Echo(args []string) string {
	s := strings.Join(args, " ")
	return s
}

func EchoIdxArgs(args []string ) string {
	s := ""
	for idx, arg := range args {
		s += strconv.Itoa(idx) + ": " + arg + "\n"
	}
	return s
}