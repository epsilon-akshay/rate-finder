package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		c.Del(context.Background(), "12312313")
	})
}

func TestGetKey(t *testing.T) {
	t.Run("should return error when redis Get returns error", func(t *testing.T) {
		c := redis.NewClient(&redis.Options{
			Addr:     "localhost:6371", //wrong port to emulate error
			Password: "",
			DB:       0,
		})

		rdb := GoRedis{
			Client: c,
		}
		_, err := rdb.GetKey(context.Background(), "12312313")
		expectedErr := "dial tcp [::1]:6371: connect: connection refused"
		assert.Equal(t, err.Error(), expectedErr)
	})

	t.Run("should return no error when redis Get the value", func(t *testing.T) {
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

		res, err := rdb.GetKey(context.Background(), "12312313")
		require.NoError(t, err)
		assert.Equal(t, res, "Active")

		c.Del(context.Background(), "12312313")
	})
}
