package main

import (
	"fmt"
)

func isValidBrackets(s string) bool {
	if len(s) < 1 || len(s) > 4096 {
		return false
	}

	pasangan := map[rune]rune{
		'>': '<',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, char := range s {
		if char == '<' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == '>' || char == '}' || char == ']' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != pasangan[char] {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println("Uji Coba TRUE")
	fmt.Println(isValidBrackets("{{[<>[{{}}]]}}"))

	fmt.Println("\nUji Coba FALSE")
	fmt.Println(isValidBrackets("]"))
	fmt.Println(isValidBrackets("]["))
	fmt.Println(isValidBrackets("[>]"))
	fmt.Println(isValidBrackets("[{<]}>"))
}
