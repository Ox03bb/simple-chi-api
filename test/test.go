package test

import (
	"errors"
	"fmt"
)

func addAndMultiply(a, b int) (int, int, string, error) {
	if a < 0 || b < 0 {
		return 0, 0, "", errors.New("inputs must be non-negative")
	}
	return a + b, a * b, "Hello, World!", nil
}

func main() {
	sum, product, _, err := addAndMultiply(3, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Sum:", sum)
	fmt.Print("Product:", product, "aaaaaaa")

}
