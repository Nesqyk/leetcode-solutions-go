package problem

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the initial prefix
	prefix := strs[0]

	// Iterate over the remaining strings
	for _, str := range strs[1:] {
		// Trim the prefix until it matches the start of the current string
		for !strings.HasPrefix(str, prefix) {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}
