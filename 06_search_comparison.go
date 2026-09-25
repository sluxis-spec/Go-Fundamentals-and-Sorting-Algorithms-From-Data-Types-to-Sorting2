package main

import "fmt"

func LinearSearch(array []int, target int) (int, int) {
	comparison := 0
	for i := 0; i < len(array); i++ {
		comparison++
		if array[i] == target {
			return i, comparison
		}
	}
	return -1, comparison
}

func binarySearch(array []int, target int) (int, int) {
	left := 0
	right := len(array) - 1
	comparison := 0

	for left <= right {
		middle := (left + right) / 2
		comparison++

		if array[middle] == target {
			return middle, comparison
		} else if target < array[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1, comparison
}

func main() {
	studentIDs := []int{101, 105, 108, 112, 119, 125, 131, 140, 155}

	target1 := 101

	linearIndex, linearComparison := LinearSearch(studentIDs, target1)
	binaryIndex, binaryComparison := binarySearch(studentIDs, target1)

	fmt.Println("Searching for:", target1)
	fmt.Println("Linear Search:")
	fmt.Println("Index:", linearIndex)
	fmt.Println("Comparisons:", linearComparison)

	fmt.Println("Binary Search:")
	fmt.Println("Index:", binaryIndex)
	fmt.Println("Comparisons:", binaryComparison)

	fmt.Println()

	target2 := 131

	linearIndex, linearComparison = LinearSearch(studentIDs, target2)
	binaryIndex, binaryComparison = binarySearch(studentIDs, target2)

	fmt.Println("Searching for:", target2)

	fmt.Println("Linear Search:")
	fmt.Println("Index:", linearIndex)
	fmt.Println("Comparisons:", linearComparison)

	fmt.Println("Binary Search:")
	fmt.Println("Index:", binaryIndex)
	fmt.Println("Comparisons:", binaryComparison)
}
