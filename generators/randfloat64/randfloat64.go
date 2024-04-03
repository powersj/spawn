package randfloat64

import "math/rand"

type Randomfloat64 struct {
	Min   int64    `toml:"min"`
	Max   int64    `toml:"max"`
	Debug bool     `toml:"debug"`
	Foo   string   `toml:"foo"`
	Blah  float64  `toml:"blah"`
	Array []string `toml:"array"`
}

func (r *Randomfloat64) Generate() any {
	return rand.Float64()
}
