package main

import (
	"fmt"

	"example.com/pricing-calculator/filemanager"
	"example.com/pricing-calculator/prices"
)

func main() {
	fmt.Println("hello nimissss...!! welcome to price calculator..")
	//prices := []float64{10, 20, 30}
	var taxes []float64

	taxes = []float64{0, 0.7, 0.10, 0.15}
	doneChan := make([]chan bool, len(taxes))
	errorChan := make([]chan error, len(taxes))

	//fmt.Println(taxes)

	var result map[float64][]float64 = make(map[float64][]float64)
	_ = result
	//fmt.Println(result)

	for index, taxRate := range taxes {
		doneChan[index] = make(chan bool)
		errorChan[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		//err := priceJob.Process()
		go priceJob.Process(doneChan[index], errorChan[index])

		/* 		if err != nil {
			fmt.Println("failed to proces jobs")
			fmt.Println(err)
		} */

	}

	for index, _ := range taxes {
		select {
		case err := <-errorChan[index]:
			if err != nil {
				fmt.Println(err)
			}

		case <-doneChan[index]:
			fmt.Println("done")
		}
	}
	/*
		 	for _, doneCh := range doneChan {
				<-doneCh
			}
	*/
}
