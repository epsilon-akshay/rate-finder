package currency_rate

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetTargetConversionRateCLient(t *testing.T) {
	t.Run("should return error when request cannot be created", func(t *testing.T) {
		client := FixerClient{
			Url: "http://[::1]:namedport",
		}

		val, err := client.GetTargetConversionRate("USD")

		actualVal := float64(0)
		expectedErr := "parse \"http://[::1]:namedport\": invalid port \":namedport\" after host"

		assert.Equal(t, err.Error(), expectedErr)
		assert.Equal(t, val, actualVal)
	})

	t.Run("should return error when a non 2xx response is returned", func(t *testing.T) {
		mockHttpClient := func() (*http.Response, error) {
			stringReader := strings.NewReader("")
			stringReadCloser := io.NopCloser(stringReader)

			return &http.Response{StatusCode: 400, Body: stringReadCloser}, nil
		}

		client := FixerClient{
			Url:        "http://fixer",
			AccessKey:  "RANDOM",
			httpClient: mockHttpDoer(mockHttpClient),
		}

		val, err := client.GetTargetConversionRate("USD")

		actualVal := float64(0)
		expectedErr := "could not fetch response with statusCode 400"

		assert.Equal(t, err.Error(), expectedErr)
		assert.Equal(t, val, actualVal)
	})
}
