package q2

import (
	"strings"
)

func Palindrome(data string) (status bool) {
	text := strings.ToLower(data)
	status = true
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			status = false
		}
	}
	return
}
