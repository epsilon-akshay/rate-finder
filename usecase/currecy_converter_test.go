package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConvertedAmountFrom(t *testing.T) {
	t.Run("should return error when rate finder returns error", func(t *testing.T) {
		client := mockrateFinder(func() (float64, error) {
			return 0, errors.New("failed to convert")
		})
		service := ConversionCalculator{
			RateFinder: client,
		}

		_, err := service.GetConvertedAmountFrom("USD")
		expectedErr := "failed to convert"
		assert.Equal(t, err.Error(), expectedErr)
	})
}
