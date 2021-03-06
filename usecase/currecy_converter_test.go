package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetConvertedAmountFrom(t *testing.T) {
	t.Run("should return error when api key is not active", func(t *testing.T) {
		client := mockrateFinder(func(ctx context.Context) (float64, error) {
			return 1.2, nil
		})

		mockRepo := mockRepo{err: nil, val: "not"}

		service := ConversionCalculator{
			RateFinder:  client,
			ApiKeyStore: mockRepo,
		}

		_, err := service.GetConvertedAmountFrom(context.Background(), "USD", float64(1), "")
		expectedErr := "not found api key"
		assert.Equal(t, err.Error(), expectedErr)
	})

	t.Run("should return error when api key is not present", func(t *testing.T) {
		client := mockrateFinder(func(ctx context.Context) (float64, error) {
			return 1.2, nil
		})

		mockRepo := mockRepo{err: errors.New("api key not found")}

		service := ConversionCalculator{
			RateFinder:  client,
			ApiKeyStore: mockRepo,
		}

		_, err := service.GetConvertedAmountFrom(context.Background(), "USD", float64(1), "")
		expectedErr := "invalid api key"
		assert.Equal(t, err.Error(), expectedErr)
	})

	t.Run("should return error when rate finder returns error", func(t *testing.T) {
		client := mockrateFinder(func(ctx context.Context) (float64, error) {
			return 0, errors.New("failed to convert")
		})

		mockRepo := mockRepo{val: "Active", err: nil}

		service := ConversionCalculator{
			RateFinder:  client,
			ApiKeyStore: mockRepo,
		}

		_, err := service.GetConvertedAmountFrom(context.Background(), "USD", float64(1), "")
		expectedErr := "failed to convert"
		assert.Equal(t, err.Error(), expectedErr)
	})

	t.Run("should return rate*amount when rate finder succeeds in calculating rate", func(t *testing.T) {
		client := mockrateFinder(func(ctx context.Context) (float64, error) {
			return 1.2, nil
		})
		mockRepo := mockRepo{val: "Active", err: nil}
		service := ConversionCalculator{
			RateFinder:  client,
			ApiKeyStore: mockRepo,
		}

		res, err := service.GetConvertedAmountFrom(context.Background(), "EUR", float64(1), "")
		require.NoError(t, err)
		assert.Equal(t, res, float64(1.2))
	})

	t.Run("should return amount/rate when rate finder succeeds in calculating rate", func(t *testing.T) {
		client := mockrateFinder(func(ctx context.Context) (float64, error) {
			return 1.2, nil
		})
		mockRepo := mockRepo{val: "Active", err: nil}
		service := ConversionCalculator{
			RateFinder:  client,
			ApiKeyStore: mockRepo,
		}

		res, err := service.GetConvertedAmountFrom(context.Background(), "USD", float64(120), "")
		require.NoError(t, err)
		assert.Equal(t, res, float64(100))
	})
}
