package read

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SplitIntElements(scanner *bufio.Scanner) []int {
	var result []int

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, num)
	}

	return result
}

func SplitStringElements(scanner *bufio.Scanner) []string {
	var result []string

	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)
	}

	return result
}

func SplitIntElementsByEmptyLines(scanner *bufio.Scanner) [][]int {
	var result [][]int
	var currentGroup []int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			result = append(result, currentGroup)
			currentGroup = nil
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			currentGroup = append(currentGroup, num)
		}
	}

	if len(currentGroup) > 0 {
		result = append(result, currentGroup)
	}

	return result
}

func SplitStringElementsByEmptyLines(scanner *bufio.Scanner) [][]string {
	var result [][]string
	var currentGroup []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			result = append(result, currentGroup)
			currentGroup = nil
		} else {
			currentGroup = append(currentGroup, line)
		}
	}

	if len(currentGroup) > 0 {
		result = append(result, currentGroup)
	}

	return result
}
