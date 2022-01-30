package usecase

import (
	"context"
	"currency_converter/repository"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateKeyShouldReturnError(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6371", //wrong port to emulate error
		Password: "",
		DB:       0,
	})

	rdb := repository.GoRedis{
		Client: c,
	}

	p := ProtectApiClient{
		KeyHolder: rdb,
	}
	_, err := p.GenerateKey(context.Background())
	expectedErr := "dial tcp [::1]:6371: connect: connection refused"
	assert.Equal(t, err.Error(), expectedErr)
}

func TestGenerateKeyShouldReturnRandomKey(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //wrong port to emulate error
		Password: "",
		DB:       0,
	})

	rdb := repository.GoRedis{
		Client: c,
	}

	p := ProtectApiClient{
		KeyHolder: rdb,
	}
	res, err := p.GenerateKey(context.Background())
	require.NoError(t, err)
	assert.Equal(t, res, "ABC")
}
