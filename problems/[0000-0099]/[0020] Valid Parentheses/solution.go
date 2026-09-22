package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '(', '{', '[':
			if ch == '(' {
				stack = append(stack, ')')
			} else if ch == '{' {
				stack = append(stack, '}')
			} else {
				stack = append(stack, ']')
			}
		default:
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
