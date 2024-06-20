package problem

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	var result int
	n := len(s)

	for i := 0; i < n; i++ {
		// If the current value is less than the next value, subtract it
		if i < n-1 && symbols[s[i]] < symbols[s[i+1]] {
			result -= symbols[s[i]]
		} else {
			result += symbols[s[i]]
		}
	}

	return result
}
