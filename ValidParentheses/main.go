package main

func validParenthesis(s string) bool {
	stack := []byte{}

	matching := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for index := 0; index < len(s); index++ {

		if s[index] == '(' || s[index] == '[' || s[index] == '{' {
			stack = append(stack, s[index])

		} else {
			// Bug 2 fix — check empty before popping
			if len(stack) == 0 {
				return false
			}

			// Bug 1 fix — check if top matches
			top := stack[len(stack)-1]
			if top != matching[s[index]] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
