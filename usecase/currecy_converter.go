package usecase

type ConvertionRateFetcher interface {
	GetTargetConversionRate() (float64, error)
}

type ConversionCalculator struct {
	RateFinder ConvertionRateFetcher
}

func (c ConversionCalculator) GetConvertedAmountFrom(base string, amount float64) (float64, error) {
	euroTodollarRate, err := c.RateFinder.GetTargetConversionRate()
	if err != nil {
		return 0, err
	}
	if base == "EUR" {
		return euroTodollarRate * amount, nil
	}
	return amount / euroTodollarRate, nil
}
