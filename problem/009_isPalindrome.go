package problem

func IsPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Variables to store the original number and the reversed number
	original := x
	var reverse int

	// Reversing the number
	for x != 0 {
		reverse = (reverse * 10) + (x % 10)
		x /= 10
	}

	// Check if the reversed number is the same as the original
	return original == reverse
}
