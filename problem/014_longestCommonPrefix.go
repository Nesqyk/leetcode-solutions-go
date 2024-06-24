package problem

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) (commonPrefix string, err error) {
	if len(strs) == 0 {
		return "", fmt.Errorf("input list is empty")
	}
	prefix := strs[0]              // Start with the first string as the initial prefix
	for _, str := range strs[1:] { // Iterate over the remaining strings
		for !strings.HasPrefix(str, prefix) { // Trim the prefix until it matches the start of the current string
			if len(prefix) == 0 {
				return "", fmt.Errorf("no common prefix found")
			}
			prefix = prefix[:len(prefix)-1] // trim the prefix
		}
	}
	return prefix, nil
}
