package problem

import (
	"fmt"
	rand2 "math/rand/v2"
)

func Test() {

	var nums []int
	for i := 0; i < 10; i++ {
		rand := rand2.IntN(10)
		nums = append(nums, rand)

	}
	TwoSum, err := TwoSum(nums, 9)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	fmt.Printf("Sum of \n%v :\n%d", nums, TwoSum)
	//
	//for {
	//	select {}
	//}
}
