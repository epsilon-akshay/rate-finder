package handler

type mockService func() (float64, error)

func (m mockService) GetConvertedAmountFrom(base string, amount float64) (float64, error) {
	return m()
}
