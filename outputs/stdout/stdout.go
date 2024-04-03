package stdout

import "os"

type Stdout struct {
	Serializers []string `toml:"serializers"`
}

func (s *Stdout) GetSerializers() []string {
	return s.Serializers
}

func (s *Stdout) Write(p []byte) error {
	_, err := os.Stdout.Write(p)
	return err
}
