package service

func ExampleFunc(a, b int, op string) int {
	switch op {
	case "add":
		return a + b
	case "sub":
		return a - b
	}
	return 0
}
