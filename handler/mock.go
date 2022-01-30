package handler

import "context"

type mockService func(ctx context.Context) (float64, error)

type mockKeyGen func(ctx context.Context) (string, error)

func (m mockService) GetConvertedAmountFrom(ctx context.Context, base string, amount float64, key string) (float64, error) {
	return m(ctx)
}

func (m mockKeyGen) GenerateKey(ctx context.Context) (string, error) {
	return m(ctx)
}
