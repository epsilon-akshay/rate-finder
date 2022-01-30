package usecase

type mockrateFinder func() (float64, error)

func (m mockrateFinder) GetTargetConversionRate(base string) (float64, error) {
	return m()
}
