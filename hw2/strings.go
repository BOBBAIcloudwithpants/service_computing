package hw2

import "strings"

func PackInFields(input string) []string {
	output := strings.Fields(input)
	return output
}

func SwitchToTitle(input string) string {
	output := strings.Title(input)
	return output
}