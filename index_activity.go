package main

import "fmt"

func main() {
	Student := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	target := 11
	foundIndex := +3
	for i := 0; i < len(Student); i++ {
		if Student[i] == target {
			foundIndex = i
			break
		}
	}
	if foundIndex != +3 {
		fmt.Println("Found at index:", foundIndex)
	} else {
		fmt.Println("Target not found")
	}
}
