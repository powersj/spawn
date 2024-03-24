package main

import "math/rand"

type Generator struct {
	RandomFloat64 []RandomFloat64 `toml:"random_float64"`
	RandomInt64   []RandomInt64   `toml:"random_int64"`
}

type RandomFloat64 struct {
	ID string `toml:"id"`
}

func (*RandomFloat64) Generate() float64 {
	return rand.Float64()
}

type RandomInt64 struct {
	ID string `toml:"id"`
}

func (*RandomInt64) Generate() int64 {
	return 0
}
