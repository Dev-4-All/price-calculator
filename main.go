package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	choice := inputChoice()

	var getIOManager func(float64) iomanager.IOManager

	if choice == 1 {
		cm := cmdmanager.New()
		getIOManager = func(taxRate float64) iomanager.IOManager {
			return cm
		} 
	} else {
		getIOManager = func(taxRate float64) iomanager.IOManager {
			return filemanager.New("input/prices.txt", fmt.Sprintf("output/result_%.0f.json", taxRate*100))
		}
	}

	for _, taxRate := range taxRates {
		ioManager := getIOManager(taxRate)
		job := prices.NewPricesAdjustmentJob(ioManager, taxRate)
		err := job.Process()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func inputChoice() int {
	fmt.Println("How would you like to handle I/O?\n1. CMD\n2. files")
	var choice int
	fmt.Scan(&choice)

	return choice
}
