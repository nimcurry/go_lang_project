package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)

	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubledNumber := transformNumbers(&numbers, double)
	tripledNumber := transformNumbers(&numbers, triple)

	fmt.Println("double number: ", doubledNumber)
	fmt.Println("triple number: ", tripledNumber)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Every anaonymus function is closure. if you use a variable from scope in which function is created, so from an outer scope like factor
// here, then the value of that outer scope variable or parameter is locked in into this created function
// In this case for eg in main function calling doubled or tripled there by passing separate values. this keeps the number passed as constant
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
