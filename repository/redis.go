package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type GoRedis struct {
	Client *redis.Client
}

func (r GoRedis) SetKey(ctx context.Context, key string) error {
	err := r.Client.Set(ctx, key, "Active", 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r GoRedis) GetKey(ctx context.Context, key string) (string, error) {
	res, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}
