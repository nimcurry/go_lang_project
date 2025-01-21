package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	var prices []string
	fmt.Println("Please enter your prices. Confirm every price with Enter")
	for {
		var price string
		fmt.Println("Price:")
		fmt.Scan(&price)

		if price == "0" {
			break
			//os.Exit(0)
		}
		prices = append(prices, price)
	}
	return prices, nil

}

func (fm CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)

	return nil
}

func New() CMDManager {
	return CMDManager{}
}
