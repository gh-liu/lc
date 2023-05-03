package main

func main() {
}

func isValid(s string) bool {
	stack := []byte{}
	for i := range s {
		stack = append(stack, s[i])
		if len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}
