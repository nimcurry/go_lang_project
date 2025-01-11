package main

import "fmt"

/* type Product struct {
	title string
	id    string
	price float64
} */

/* func main() {
	var productNames [4]string
	productNames = [4]string{"A book"}
	prices := [4]float64{29.99, 13.99, 10.00, 45.0}
	fmt.Println(prices)
	productNames[1] = "a carpet"
	fmt.Println(productNames[1])

	featuredPrices := prices[1:3]

	featuredPrice1s := prices[:3] //all from the begining

	featuredPrice2s := prices[2:] //all till the end

	finalFeaturedPries := featuredPrice2s
	finalFeaturedPries[0] = 199.99

	fmt.Print(featuredPrices, featuredPrice1s, featuredPrice2s)

	fmt.Println("highlighted prices", finalFeaturedPries)

	//length provides the total number of items in sliced array
	//cap provides the total capacity of slice which in turn points to the total length of original array
	//also understand that cap starts from left to right. so if slice starts from 1 till end the total cap will be total size -1
	//which means cap can traverse or select more towards end or right of the array but not to the left.
	fmt.Println(len(featuredPrice1s), cap(featuredPrice1s))
}
*/

func main() {
	prices := []float64{10.00, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	//prices[2] = 11.99

	updatedPrices := append(prices, 5.99)
	prices = append(prices, 5.99)

	prices = prices[1:]
	fmt.Println(prices, updatedPrices)

	//The append function works by appending one or more individual elements to a slice, not an entire slice as a single element. hence use "..."
	discountPrices := []float64{0.99, 8.99, 20.99}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
