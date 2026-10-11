package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		ioManager := filemanager.New("input/prices.txt", fmt.Sprintf("output/result_%.0f.json", taxRate*100))
		job := prices.NewPricesAdjustmentJob(*ioManager, taxRate)
		err := job.Process()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
