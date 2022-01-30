package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetKey(t *testing.T) {
	t.Run("should return error when redis returns error", func(t *testing.T) {
		c := redis.NewClient(&redis.Options{
			Addr:     "localhost:6371", //wrong port to emulate error
			Password: "",
			DB:       0,
		})

		rdb := GoRedis{
			Client: c,
		}
		err := rdb.SetKey(context.Background(), "12312313")
		expectedErr := "dial tcp [::1]:6371: connect: connection refused"
		assert.Equal(t, err.Error(), expectedErr)
	})

	t.Run("should return no error when redis sets the value", func(t *testing.T) {
		c := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", //wrong port to emulate error
			Password: "",
			DB:       0,
		})

		rdb := GoRedis{
			Client: c,
		}

		err := rdb.SetKey(context.Background(), "12312313")
		assert.Equal(t, err, nil)
	})
}
