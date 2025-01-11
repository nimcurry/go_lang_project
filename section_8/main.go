package main

import "fmt"

func main() {
	numbers := []int{10, 15, 20}
	sum := sumup(1, 10, 15, 25)
	//sum := sumup(numbers)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum = sum + val
	}

	return sum
}
