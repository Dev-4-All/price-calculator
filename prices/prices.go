package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type PricesAdjustmentJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxAdjustedPrices map[string]string
}

func NewPricesAdjustmentJob(taxRate float64) *PricesAdjustmentJob {
	return &PricesAdjustmentJob{
		TaxRate: taxRate,
	}
}

func (job *PricesAdjustmentJob) loadData() error {
	prices, err := filemanager.ReadLines("input/prices.txt")

	if err != nil {
		return err
	}

	job.InputPrices, err = conversion.StringsToFloats(prices)

	if err != nil {
		return err
	}

	return nil
}

func (job *PricesAdjustmentJob) Process() error {
	err := job.loadData()

	if err != nil {
		return err
	}

	adjustedPrices := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		adjustedPrice := price * (1 + job.TaxRate)
		adjustedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", adjustedPrice)
	}

	job.TaxAdjustedPrices = adjustedPrices

	err = filemanager.WriteJSON(fmt.Sprintf("output/result_%.0f.json", job.TaxRate*100), job)

	if err != nil {
		return err
	}

	return nil
}
