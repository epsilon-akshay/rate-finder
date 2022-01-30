package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		_, err := service.GetConvertedAmountFrom("USD", float64(1))
		expectedErr := "failed to convert"
		assert.Equal(t, err.Error(), expectedErr)
	})

	t.Run("should return rate*amount when rate finder succeeds in calculating rate", func(t *testing.T) {
		client := mockrateFinder(func() (float64, error) {
			return 1.2, nil
		})
		service := ConversionCalculator{
			RateFinder: client,
		}

		res, err := service.GetConvertedAmountFrom("USD", float64(1))
		require.NoError(t, err)
		assert.Equal(t, res, float64(1.2))
	})
}
