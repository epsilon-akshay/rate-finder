package usecase

import "context"

type Repository interface {
	SetKey(ctx context.Context, key string) error
}

type ProtectApiClient struct {
	KeyHolder Repository
}

func (p ProtectApiClient) GenerateKey(ctx context.Context) (string, error) {
	randomKey := "ABC"
	err := p.KeyHolder.SetKey(ctx, randomKey)
	if err != nil {
		return "", err
	}
	return randomKey, nil
}
