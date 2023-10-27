package leetcode

import (
	"github.com/luo/leetcode-kit/util"
)

// code your solution here

// leetcode 1
func twoSum(nums []int, target int) []int {
    for i, x := range nums {
        for j := i + 1; j < len(nums); j++ {
            if x+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

// use the kit
func CompareIntSlice(a, b []int) bool {
	return util.IntSliceIsEqual(a, b)
}