package currency_rate

import "net/http"

type mockHttpDoer func() (*http.Response, error)

func (m mockHttpDoer) Do(req *http.Request) (*http.Response, error) {
	return m()
}
