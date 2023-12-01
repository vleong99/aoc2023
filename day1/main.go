package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ans := addIntSlices(read.splitIntElementsByEmptyLines(scanner))

	fmt.Println(ans[len(ans)-1] + ans[len(ans)-2] + ans[len(ans)-3])

}

func addIntSlices(nums [][]int) []int {
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
