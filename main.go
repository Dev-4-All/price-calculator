package main

import (
	"fmt"

	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		job := prices.NewPricesAdjustmentJob(taxRate)
		err := job.Process()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
