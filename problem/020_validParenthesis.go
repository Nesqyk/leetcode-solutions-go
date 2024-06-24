package problem

func isValid(s string) bool {
	var stack []rune
	matchingBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			// If the character is an opening bracket, push it onto the stack
			stack = append(stack, char)
		case ')', ']', '}':
			// If the character is a closing bracket, check if it matches the top of the stack
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			// Pop the top element from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, all brackets were properly closed
	return len(stack) == 0
}
