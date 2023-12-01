package slicesConv

import "slices"

func AddIntSlices(nums [][]int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		var ans int
		for j := 0; j < len(nums[i]); j++ {
			ans += nums[i][j]
		}
		result = append(result, ans)
	}

	slices.Sort(result)
	return result
}
