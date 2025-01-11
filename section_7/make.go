package main

import "fmt"

func main() {
	//slice is just an window in array.
	//pre allocation could be done using "make function"  in go array. this avoids creation and recreation of arrays

	userNames := make([]string, 2, 5)

	//userNames := []string{}
	userNames[0] = "Julie"
	userNames = append(userNames, "nimish")
	userNames = append(userNames, "navjot")

	fmt.Println(userNames)

	for key, values := range userNames {
		fmt.Println("key", key)
		fmt.Println("values", values)
	}
}
