package currency_rate

import (
	"net/http"
)

type FixerClient struct {
	AccessKey string
	Url       string
}

func (client FixerClient) GetTargetConversionRate(base string) (float64, error) {
	_, err := http.NewRequest("GET", client.Url, nil)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
