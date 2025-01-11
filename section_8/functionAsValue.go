package main

import "fmt"

type transformN func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{3, 1, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFunc1 := getTransformerFunction(&numbers)
	transformerFunc2 := getTransformerFunction(&moreNumbers)

	transformedNum1 := transformNumbers(&numbers, transformerFunc1)
	transformedNum2 := transformNumbers(&moreNumbers, transformerFunc2)

	fmt.Println("num1: ", transformedNum1)
	fmt.Println("num2: ", transformedNum2)

}

func transformNumbers(numbers *[]int, transform transformN) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformN {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
