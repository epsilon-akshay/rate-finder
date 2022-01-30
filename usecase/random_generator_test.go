package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandStringRunes(t *testing.T) {
	r := RandomGen{
		Size:        4,
		LetterRunes: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}
	assert.Equal(t, len(r.RandStringRunes()), 4)
}
