package main

func makeGood(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1]^32 == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
