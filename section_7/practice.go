package main

import "fmt"

func main() {
	//1
	hobbies := [3]string{"fukking", "thukking", "looking"}

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	mainHobbies := hobbies[0:2]

	fmt.Println(mainHobbies)

	//4
	fmt.Println(cap(mainHobbies))
	//reslicing on the existing slicing would require you to specify the end index as well
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5
	courseGoals := []string{"Learn go!", "Learn all the contents of go"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "learn all the details"
	courseGoals = append(courseGoals, "learn all the basics")
	fmt.Println(courseGoals)

	//7
	type product struct {
		id          int
		productName string
	}
	products := []product{
		{
			1,
			"first-product"},
		{
			2,
			"first-product"},
	}
	fmt.Println(products)
	newProduct := product{3, "lavdijot"}
	products = append(products, newProduct)

	fmt.Println(products)
}
