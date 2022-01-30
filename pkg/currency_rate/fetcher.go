package currency_rate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}
type FixerClient struct {
	AccessKey  string
	Url        string
	HttpClient Doer
}

type FixerResponse struct {
	Success bool   `json:"success"`
	Base    string `json:"base"`
	Rates   struct {
		USD float64 `json:"USD"`
		EUR float64 `json:"EUR"`
	} `json:"rates"`
}

func (c FixerClient) GetTargetConversionRate() (float64, error) {
	req, err := http.NewRequest("GET", c.Url, nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("access_key", c.AccessKey)
	req.URL.RawQuery = q.Encode()

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	httpRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("could not fetch response with statusCode %v", res.StatusCode)
	}

	var fixerRes FixerResponse

	err = json.Unmarshal(httpRes, &fixerRes)
	if err != nil {
		return 0, err
	}

	return fixerRes.Rates.USD, nil

}
