package problem

import "fmt"

func IsPalindrome(x int) (check bool, err error) {
	if x < 0 {
		return false, fmt.Errorf("negative integers are not palindromes")
	}
	original := x
	var reverse int

	// Reversing the number
	for x != 0 {
		reverse = (reverse * 10) + (x % 10)
		x /= 10
	}

	// Check if the reversed number is the same as the original
	return original == reverse, nil
}
