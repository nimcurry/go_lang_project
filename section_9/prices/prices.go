package prices

import (
	"fmt"

	"example.com/pricing-calculator/conversion"
	"example.com/pricing-calculator/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		//fmt.Println("failed reading lines")
		//fmt.Println(err)
		return err

	}
	var prices []float64
	//prices = make([]float64, len(lines))
	prices, err = conversion.StringToFloat(lines)

	if err != nil {
		//fmt.Println("failed converting string to floadt")
		//fmt.Println(err)
		return err

	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPricesJob) Process(doneChan chan bool, errorChan chan error) {
	//error
	err := job.LoadData()

	//errorChan <- errors.New("error")

	if err != nil {
		//return err
		errorChan <- err
		return
	}
	//to store tax included prices
	result := make(map[string]string)

	//fmt.Println(index, taxes)
	for _, prices := range job.InputPrices {
		taxIncludedPrices := prices * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", prices)] = fmt.Sprintf("%.2f", taxIncludedPrices)

	}
	job.TaxIncludedPrices = result
	//fmt.Println(result)
	//return job.IOManager.WriteResult(job)
	job.IOManager.WriteResult(job)
	doneChan <- true
}

// constructor function - goal to return taxIncludedPricesJob
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
