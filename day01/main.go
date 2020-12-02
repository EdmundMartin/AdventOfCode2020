package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func threeSum(seen []int, target int) int {
	// TODO - Bad time complexity but simple
	for i := 0; i < len(seen)-2; i++ {
		for j := i + 1; j < len(seen)-1; j++ {
			for k := j + 1; k < len(seen); k++ {
				sum := seen[i] + seen[j] + seen[k]
				if sum == target {
					return seen[i] * seen[j] * seen[k]
				}
			}
		}
	}
	return -1
}

func main() {
	fileHandle, _ := os.Open("day01/input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	target := 2020
	seen := make(map[int]interface{})
	all := []int{}
	for fileScanner.Scan() {
		res, _ := strconv.Atoi(fileScanner.Text())
		other := target - res
		if _, ok := seen[other]; ok {
			finalResult := res * other
			fmt.Printf("Answer is %d\n", finalResult)
		}
		seen[res] = nil
		all = append(all, res)
	}
	value := threeSum(all, 2020)
	fmt.Printf("Second Answer is %d\n", value)
}
