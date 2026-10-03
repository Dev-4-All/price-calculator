package prices

import "fmt"

type PricesAdjustmentJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxAdjustedPrices map[string]float64
}

func NewPricesAdjustmentJob(taxRate float64) *PricesAdjustmentJob {
	return &PricesAdjustmentJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10.0, 20.0, 30.0},
	}
}

func (job PricesAdjustmentJob) Process() {
	adjustedPrices := make(map[string]float64, len(job.InputPrices))

	for _, price := range job.InputPrices {
		adjustedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(adjustedPrices)
}
