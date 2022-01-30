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

type mockRepo struct {
	val string
	err error
}

func (r mockRepo) SetKey(ctx context.Context, key string) error {
	return r.err
}

func (r mockRepo) GetKey(ctx context.Context, key string) (string, error) {
	return r.val, r.err
}
