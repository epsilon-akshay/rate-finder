package usecase

import (
	"context"
	"errors"
)

type ConvertionRateFetcher interface {
	GetTargetConversionRate(ctx context.Context) (float64, error)
}

type ConversionCalculator struct {
	RateFinder  ConvertionRateFetcher
	ApiKeyStore Repository
}

func (c ConversionCalculator) GetConvertedAmountFrom(ctx context.Context, base string, amount float64, key string) (float64, error) {
	res, err := c.ApiKeyStore.GetKey(ctx, key)
	if err != nil {
		return 0, errors.New("invalid api key")
	}

	if res != "Active" {
		return 0, errors.New("not found api key")
	}

	euroTodollarRate, err := c.RateFinder.GetTargetConversionRate(ctx)
	if err != nil {
		return 0, err
	}
	if base == "EUR" {
		return euroTodollarRate * amount, nil
	}
	return amount / euroTodollarRate, nil
}
