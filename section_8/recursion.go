package main

import "fmt"

func main() {
	fact := factorial(8)
	fmt.Println(fact)

}

func factorial(number int) int {
	if number == 0 {

		return 1
	} else {
		return number * factorial(number-1)
	}
	/*
		 	result := 1
			//i := 1
			for i := 1; i <= number; i++ {
				result = result * i

			}

			return result
	*/
}

/*
factorial(5) -> executes -> 5 * factorial(4)
factorial(4) -> executes -> 4 * factorial(3)
factorial(3) -> executes -> 3 * factorial(2)
factorial(2) -> executes -> 2 * factorial(1)
factorial(0) -> number is 0 executes -> in the if condition exits with value 1
Now 1 -> 2*1 -> 3*2 -> 6*4 -> 24*5 -> 120
*/
