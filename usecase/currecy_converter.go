package usecase

import "context"

type ConvertionRateFetcher interface {
	GetTargetConversionRate(ctx context.Context) (float64, error)
}

type ConversionCalculator struct {
	RateFinder ConvertionRateFetcher
}

func (c ConversionCalculator) GetConvertedAmountFrom(ctx context.Context, base string, amount float64) (float64, error) {
	euroTodollarRate, err := c.RateFinder.GetTargetConversionRate(ctx)
	if err != nil {
		return 0, err
	}
	if base == "EUR" {
		return euroTodollarRate * amount, nil
	}
	return amount / euroTodollarRate, nil
}
