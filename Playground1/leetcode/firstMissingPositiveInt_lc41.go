package leetcode

import (
	"math"
)

func firstMissingPositive(nums []int) int {
	for i:=0;i<len(nums);i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 2 ;
		}
	}

	for i:=0; i<len(nums);i++ {
		ele := int(math.Abs(float64(nums[i])));
		if ele >= 1 && float64(ele) <= float64(len(nums)){
			index := ele -1;
			nums[index] = -1 * int(math.Abs(float64(nums[index])));
		}
	}
	for i:=0;i<len(nums);i++ {
		if nums[i]>0 {
			return i+1;
		}
	}
	return len(nums) + 1; 
}
