package usecase

import "context"

type mockrateFinder func(ctx context.Context) (float64, error)

func (m mockrateFinder) GetTargetConversionRate(ctx context.Context) (float64, error) {
	return m(ctx)
}

type mockKeyGen func() string

func (m mockKeyGen) RandStringRunes() string {
	return m()
}
