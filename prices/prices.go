package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type PricesAdjustmentJob struct {
	ioManager         iomanager.IOManager
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxAdjustedPrices map[string]string `json:"tax_adjusted_prices"`
}

func NewPricesAdjustmentJob(ioManager iomanager.IOManager, taxRate float64) *PricesAdjustmentJob {
	return &PricesAdjustmentJob{
		ioManager: ioManager,
		TaxRate:   taxRate,
	}
}

func (job *PricesAdjustmentJob) loadData() error {
	prices, err := job.ioManager.ReadLines()

	if err != nil {
		return err
	}

	job.InputPrices, err = conversion.StringsToFloats(prices)

	if err != nil {
		return err
	}

	return nil
}

func (job *PricesAdjustmentJob) Process(doneChan chan bool, errChan chan error) {
	err := job.loadData()

	if err != nil {
		errChan <- err
		return
	}

	adjustedPrices := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		adjustedPrice := price * (1 + job.TaxRate)
		adjustedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", adjustedPrice)
	}

	job.TaxAdjustedPrices = adjustedPrices

	err = job.ioManager.WriteResults(job)

	if err != nil {
		errChan <- err
		return
	}

	doneChan <- true
}
