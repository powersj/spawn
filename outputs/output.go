package outputs

type Output struct {
	Stdout []Stdout `toml:"stdout"`
	Stderr []Stderr `toml:"stderr"`
	File   []File   `toml:"file"`
}

type Stdout struct {
	Serializers []string `toml:"serializers"`
}

type Stderr struct {
	Serializers []string `toml:"serializers"`
}

type File struct {
	Serializers []string `toml:"serializers"`
	Filename    string   `toml:"Filename"`
}
