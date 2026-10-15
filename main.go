package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		errChans[idx] = make(chan error)

		ioManager := filemanager.New("input/prices.txt", fmt.Sprintf("output/result_%.0f.json", taxRate*100))
		job := prices.NewPricesAdjustmentJob(ioManager, taxRate)
		go job.Process(doneChans[idx], errChans[idx])
	}

	for idx := range taxRates {
		select {
		case <-doneChans[idx]:
			fmt.Println("Done!")
		case err := <-errChans[idx]:
			fmt.Println(err)
		}
	}
}
