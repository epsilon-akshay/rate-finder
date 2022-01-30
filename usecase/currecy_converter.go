package usecase

type ConvertionRateFetcher interface {
	GetTargetConversionRate(base string) (float64, error)
}

type ConversionCalculator struct {
	RateFinder ConvertionRateFetcher
}

func (c ConversionCalculator) GetConvertedAmountFrom(base string) (float64, error) {
	_, err := c.RateFinder.GetTargetConversionRate(base)
	return 0, err
}
