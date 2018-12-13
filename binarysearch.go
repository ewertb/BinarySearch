package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getMiddle(start int, end int) int {
	return start + (end-start)/2
}

func search(numbers []int, find int) bool {
	sort.Ints(numbers)
	var start = 0
	var end = len(numbers) - 1
	var mid = getMiddle(start, end)
	var found = false

	for found == false {
		if end < start {
			break
		}

		if numbers[mid] < find {
			start = mid + 1
			mid = getMiddle(start, end)
			continue
		}

		if numbers[mid] > find {
			end = mid - 1
			mid = getMiddle(start, end)
			continue
		}

		if numbers[mid] == find {
			found = true
		}
	}

	return found
}

func main() {
	var numbers = []int{1, 3, 9, 7, 5, 4, 6, 2, 8}
	var find, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(search(numbers, find))
}
