package generators

import "math/rand"

var Registry = map[string]func() Generator{
	"randomfloat64": func() Generator { return &Randomfloat64{} },
	"randomint64":   func() Generator { return &Randomint64{} },
}

type Generator interface {
	Generate() any
}

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

type Randomint64 struct {
	Min int64 `toml:"min"`
	Max int64 `toml:"max"`
}

func (r *Randomint64) Generate() any {
	return rand.Int63()
}
