package app

import "strings"

// Baset give you diff sets of genotype
func Baset(line string) string {
	slice := strings.Split(line, "")
	var out string
	if slice[0] == slice[1] {
		out = line
	} else {
		out = slice[0] + slice[1] + ";" + slice[1] + slice[0]
	}
	return out
}
