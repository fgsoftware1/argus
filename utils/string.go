package utils

import "strings"

func SplitLines(stringSet string) []string {
	return strings.Split(stringSet, "\n")
}

func CenterString(s string, width int) string {
	lines := strings.Split(s, "\n")
	var centeredLines []string

	for _, line := range lines {
		lineLength := len([]rune(line))
		padding := (width - lineLength) / 2

		if padding < 0 {
			padding = 0
		}

		centeredLines = append(centeredLines, strings.Repeat(" ", padding)+line)
	}

	return strings.Join(centeredLines, " ")
}
