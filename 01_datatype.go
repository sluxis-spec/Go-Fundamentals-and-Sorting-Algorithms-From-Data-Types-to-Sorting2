package main

import "fmt"

func main() {
	var name string = "Danish"
	var age int = 18
	var gpa float64 = 3.75
	var active bool = true

	fmt.Println("Name =", name)
	fmt.Println("Age =", age)
	fmt.Println("GPA =", gpa)
	fmt.Println("Active =", active)

	Student := []string{"Danish,", "Udin,", "Dimas,", "Mesya,", "Raul,"}
	target := "Danish,"
	foundIndex := -1
	for i := 0; i < len(Student); i++ {
		if Student[i] == target {
			foundIndex = i
			break
		}
	}
	if foundIndex != -1 {
		fmt.Println("Found at index:", foundIndex)
	} else {
		fmt.Println("Target not found")
	}
}
