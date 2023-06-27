package main

import (
	"fmt"
	"math/rand"
	"time"
)

func make_random_slice(num_items, max int) []int {
	var rand_slice []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_items; i++ {
		rand_slice = append(rand_slice, rand.Intn(max))
	}

	return rand_slice
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func print_slice(arr []int, num_items int) {
	num_elems := Min(len(arr), num_items)
	fmt.Println(arr[:num_elems])
}

func linear_search(values []int, target int) (index, num_tests int) {
	for i, _ := range values {
		if values[i] == target {
			return i, i + 1
		}
	}
	return -1, len(values)
}

func main() {
	var num_items, max, target int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	values := make_random_slice(num_items, max)
	print_slice(values, 40)
	fmt.Println()

	fmt.Printf("Target: ")
	fmt.Scanln(&target)

	i, v := linear_search(values, target)
	if i == -1 {
		fmt.Printf("Target %d not found, %d tests", target, v)
	} else {
		fmt.Printf("values[%d] = %d, %d tests", i, target, v)
	}
}
