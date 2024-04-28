package countsmileys

import "strings"

func CountSmileys(inputString []string) int {
	count := 0

	for _, sign := range inputString {
		if len(sign) == 2 {
			if findSignOfSmile(sign) {
				count++
			}
		} else if len(sign) == 3 {
			if findSignOfSmile(sign) && strings.ContainsAny(sign[1:2], "-~") {
				count++
			}
		}
	}

	return count
}

func findSignOfSmile(sign string) bool {
	result := strings.ContainsAny(sign, ":;") && strings.ContainsAny(sign, ")D")
	return result
}
