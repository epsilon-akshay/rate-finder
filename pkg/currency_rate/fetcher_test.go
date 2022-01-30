package currency_rate

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			HttpClient: mockHttpDoer(mockHttpClient),
		}

		val, err := client.GetTargetConversionRate("USD")

		actualVal := float64(0)
		expectedErr := "could not fetch response with statusCode 400"

		assert.Equal(t, err.Error(), expectedErr)
		assert.Equal(t, val, actualVal)
	})

	t.Run("should return error when a client.do throws errors", func(t *testing.T) {
		mockHttpClient := func() (*http.Response, error) {
			return nil, errors.New("random error")
		}

		client := FixerClient{
			Url:        "http://fixer",
			AccessKey:  "RANDOM",
			HttpClient: mockHttpDoer(mockHttpClient),
		}

		val, err := client.GetTargetConversionRate("USD")

		actualVal := float64(0)
		expectedErr := "random error"

		assert.Equal(t, err.Error(), expectedErr)
		assert.Equal(t, val, actualVal)
	})

	t.Run("should return appropriate target rate for Euro when a client.do returns success", func(t *testing.T) {
		mockHttpClient := func() (*http.Response, error) {
			stringReader := strings.NewReader(`{"rates": {"EUR": 1.1}}`)
			stringReadCloser := io.NopCloser(stringReader)

			return &http.Response{StatusCode: 200, Body: stringReadCloser}, nil
		}

		client := FixerClient{
			Url:        "http://fixer",
			AccessKey:  "RANDOM",
			HttpClient: mockHttpDoer(mockHttpClient),
		}

		val, err := client.GetTargetConversionRate("USD")
		require.NoError(t, err, "no error")

		actualVal := float64(1.1)
		assert.Equal(t, val, actualVal)
	})

	t.Run("should return appropriate target rate for Dollar when a client.do returns success", func(t *testing.T) {
		mockHttpClient := func() (*http.Response, error) {
			stringReader := strings.NewReader(`{"rates": {"USD": 1.1}}`)
			stringReadCloser := io.NopCloser(stringReader)

			return &http.Response{StatusCode: 200, Body: stringReadCloser}, nil
		}

		client := FixerClient{
			Url:        "http://fixer",
			AccessKey:  "RANDOM",
			HttpClient: mockHttpDoer(mockHttpClient),
		}

		val, err := client.GetTargetConversionRate("EUR")
		require.NoError(t, err, "no error")

		actualVal := float64(1.1)
		assert.Equal(t, val, actualVal)
	})

	t.Run("should return error when unmarshalling fails", func(t *testing.T) {
		mockHttpClient := func() (*http.Response, error) {
			stringReader := strings.NewReader(`{"rates": {"USD": "1.1"}}`)
			stringReadCloser := io.NopCloser(stringReader)

			return &http.Response{StatusCode: 200, Body: stringReadCloser}, nil
		}

		client := FixerClient{
			Url:        "http://fixer",
			AccessKey:  "RANDOM",
			HttpClient: mockHttpDoer(mockHttpClient),
		}

		val, err := client.GetTargetConversionRate("EUR")

		actualVal := float64(0)
		expectedErr := "json: cannot unmarshal string into Go struct field .rates.USD of type float64"

		assert.Equal(t, err.Error(), expectedErr)
		assert.Equal(t, val, actualVal)
	})

}
