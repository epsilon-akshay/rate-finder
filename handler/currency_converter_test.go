package handler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertCurrency(t *testing.T) {
	t.Run("should return 400 response when amount passed is not float", func(t *testing.T) {
		m := mockService(func() (float64, error) {
			return 64, nil
		})

		ts := httptest.NewServer(ConvertCurrency(m))
		defer ts.Close()

		url := fmt.Sprintf("%s?base=USD&amount=ooo", ts.URL)
		res, err := http.Get(url)
		require.NoError(t, err)

		expectedResponse, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		require.NoError(t, err)

		res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(expectedResponse), `{"success":false,"err":"strconv.ParseFloat: parsing \"ooo\": invalid syntax"}`)
		ConvertCurrency(m)
	})
}
