package usecase

type ConvertionRateFetcher interface {
	GetTargetConversionRate(base string) (float64, error)
}

type ConversionCalculator struct {
	RateFinder ConvertionRateFetcher
}

func (c ConversionCalculator) GetConvertedAmountFrom(base string, amount float64) (float64, error) {
	rate, err := c.RateFinder.GetTargetConversionRate(base)
	if err != nil {
		return 0, err
	}

	return rate * amount, nil
}
