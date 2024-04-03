package randint64

import "math/rand"

type Randomint64 struct {
	Min int64 `toml:"min"`
	Max int64 `toml:"max"`
}

func (r *Randomint64) Generate() any {
	return rand.Int63()
}
