package rotatestring

import "strings"

func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	} else {
		s = s + s
	}

	return strings.Contains(s, goal)
}
