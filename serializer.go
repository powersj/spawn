package main

type Serializer struct {
	Templates []Template `toml:"template"`
}

type Template struct {
	ID       string `toml:"id"`
	Template string `toml:"template"`
}
