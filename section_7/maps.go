package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
func main() {
	//type aliases for good developer experiences

	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["AWS"])
	websites["linkedin"] = "https://linkedin.com"

	fmt.Println(websites)
	delete(websites, "AWS")
	fmt.Println(websites)

	//make function is all about pre-allocating a memory. this makes map or array structures more efficient
	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.8
	courseRatings["c#"] = 3.78
	courseRatings["pythong"] = 4.5

	//fmt.Println(courseRatings)
	floatMap.output(courseRatings)

	for index, value := range websites {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

	//if you dont care for individual items or values u can also use for range websites
}
