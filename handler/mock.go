package handler

import "context"

type mockService func() (float64, error)

type mockKeyGen func(ctx context.Context) (string, error)

func (m mockService) GetConvertedAmountFrom(base string, amount float64) (float64, error) {
	return m()
}

func (m mockKeyGen) GenerateKey(ctx context.Context) (string, error) {
	return m(ctx)
}
