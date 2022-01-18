package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int
	y := 0
	var nums [10000]int

	n, _ := fmt.Scan(&x)

	for n == 1 {
		nums[y] = x
		//fmt.Printf("Array before sorting: %v\n", nums[0:y+1])
		sort.Ints(nums[0 : y+1])
		//fmt.Printf("Array after sorting: %v\n", nums[0:y+1])
		if y == 0 {
			fmt.Println(nums[0])
		} else if y%2 != 0 {
			fmt.Println((nums[y/2] + nums[(y/2)+1]) / 2)
		} else if y%2 == 0 {
			fmt.Println(nums[(y / 2)])
		}

		y++
		n, _ = fmt.Scan(&x)
	}

}
