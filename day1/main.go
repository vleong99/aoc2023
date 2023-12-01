package main

import (
	"aoc2023/utilities/read"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := read.SplitStringElements(scanner)

	result := 0
	// for i := 0; i < len(input); i++ {
	// 	line := input[i]
	// 	ans := ""
	// 	for j := 0; j < len(line); j++ {
	// 		if line[j] >= 48 && line[j] <= 57 {
	// 			ans += string(line[j])
	// 			break
	// 		}
	// 	}
	// 	for k := len(line) - 1; k >= 0; k-- {
	// 		if line[k] >= 48 && line[k] <= 57 {
	// 			ans += string(line[k])
	// 			break
	// 		}
	// 	}
	// 	val, err := strconv.Atoi(ans)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	result += val
	// }
	// fmt.Println(result)

	numbersAsText := make(map[string]string)
	numbersAsText["zero"] = "0"
	numbersAsText["one"] = "1"
	numbersAsText["two"] = "2"
	numbersAsText["three"] = "3"
	numbersAsText["four"] = "4"
	numbersAsText["five"] = "5"
	numbersAsText["six"] = "6"
	numbersAsText["seven"] = "7"
	numbersAsText["eight"] = "8"
	numbersAsText["nine"] = "9"

	for i := 0; i < len(input); i++ {
		line := input[i]

		smallest := ""
		smallestIndex := 0
		largest := ""
		largestIndex := 0

		for j := 0; j < len(line); j++ {
			if line[j] >= 48 && line[j] <= 57 {
				smallest = string(line[j])
				smallestIndex = j
				break
			}
		}
		for k := len(line) - 1; k >= 0; k-- {
			if line[k] >= 48 && line[k] <= 57 {
				largest = string(line[k])
				largestIndex = k
				break
			}
		}

		for k, v := range numbersAsText {
			first := strings.Index(line, k)
			if first >= 0 && first <= smallestIndex {
				smallestIndex = first
				smallest = v
			}

			last := strings.LastIndex(line, k)
			if last >= 0 && last >= largestIndex {
				largestIndex = last
				largest = v
			}
		}

		val, err := strconv.Atoi(smallest + largest)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)
		result += val

	}

	fmt.Println(result)

}
