package problem

import (
	"fmt"
	"github.com/alphadose/haxmap"
)

func TwoSum(nums []int, target int) (result []int, err error) {
	numMap := haxmap.New[int, int]() // Create an empty map to store numbers and their indices.
	for i, num := range nums {       // Loop over each number in the nums array with its index.
		complement := target - num                 // Calculate the complement of the current number.
		if idx, ok := numMap.Get(complement); ok { // Check if the complement is already in the map.

			return []int{idx, i}, nil // If found, return the indices of the complement and the current number.
		}
		numMap.Set(num, i) // Add the current number and its index to the map.
	}

	return []int{}, fmt.Errorf("no solution is found") // If no solution ins found, return an empty slice.
}
