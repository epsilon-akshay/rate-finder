package usecase

import "math/rand"

type RandomGen struct {
	LetterRunes []rune
	Size        int
}

func (r RandomGen) RandStringRunes() string {
	b := make([]rune, r.Size)
	for i := range b {
		b[i] = r.LetterRunes[rand.Intn(len(r.LetterRunes))]
	}
	return string(b)
}
