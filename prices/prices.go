package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type PricesAdjustmentJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxAdjustedPrices map[string]string
}

func NewPricesAdjustmentJob(ioManager filemanager.FileManager, taxRate float64) *PricesAdjustmentJob {
	return &PricesAdjustmentJob{
		IOManager: ioManager,
		TaxRate: taxRate,
	}
}

func (job *PricesAdjustmentJob) loadData() error {
	prices, err := job.IOManager.ReadLines()

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

	err = job.IOManager.WriteJSON(job)

	if err != nil {
		return err
	}

	return nil
}
