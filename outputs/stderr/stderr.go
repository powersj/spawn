package stderr

import "os"

type Stderr struct {
	Serializers []string `toml:"serializers"`
}

func (s *Stderr) GetSerializers() []string {
	return s.Serializers
}

func (s *Stderr) Write(p []byte) error {
	_, err := os.Stderr.Write(p)
	return err
}
