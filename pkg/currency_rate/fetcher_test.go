package currency_rate

import (
	"github.com/stretchr/testify/assert"
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
}
