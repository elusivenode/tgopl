package echo

import "strings"

func Echo(args []string) string {
	s := strings.Join(args, " ")
	return s
}
