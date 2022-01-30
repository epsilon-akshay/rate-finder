package usecase

type mockrateFinder func() (float64, error)

func (m mockrateFinder) GetTargetConversionRate() (float64, error) {
	return m()
}

type mockKeyGen func() string

func (m mockKeyGen) RandStringRunes() string {
	return m()
}
