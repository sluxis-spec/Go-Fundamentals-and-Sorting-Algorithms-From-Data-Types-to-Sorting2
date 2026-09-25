package main

import "fmt"

func LinearSearch(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}
	return -1
}
func main() {
	numbers := []int{12, 7, 25, 18, 9}

	fmt.Println("test 1 - exists:", LinearSearch(numbers, 18))
	fmt.Println("test 2 - not exists:", LinearSearch(numbers, 100))
}
