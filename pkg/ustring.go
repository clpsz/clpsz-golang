package pkg

import "strings"

func BreakStringByLine(str string) []string {
	return strings.Split(str, "\n")
}

