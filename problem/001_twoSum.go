package problem

import "fmt"

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Create an empty map to store numbers and their indices.
	for i, num := range nums {  // Loop over each number in the nums array with its index.
		complement := target - num             // Calculate the complement of the current number.
		if idx, ok := numMap[complement]; ok { // Check if the complement is already in the map.
			return []int{idx, i} // If found, return the indices of the complement and the current number.
		}
		numMap[num] = i // Add the current number and its index to the map.
	}

	return []int{} // If no solution is found, return an empty slice.
}

func Test() {
	fmt.Println(TwoSum([]int{2, 9, 11, 15}, 11))
	fmt.Println(IsPalindrome(121))
}
