package usecase

import "context"

type Repository interface {
	SetKey(ctx context.Context, key string) error
}

type RandomStringGenerator interface {
	RandStringRunes() string
}

type ProtectApiClient struct {
	KeyHolder    Repository
	KeyGenerator RandomStringGenerator
}

func (p ProtectApiClient) GenerateKey(ctx context.Context) (string, error) {
	randomKey := p.KeyGenerator.RandStringRunes()
	err := p.KeyHolder.SetKey(ctx, randomKey)
	if err != nil {
		return "", err
	}
	return randomKey, nil
}
