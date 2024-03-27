package outputs

import "os"

var Registry = map[string]func() Output{
	"stdout": func() Output { return &Stdout{} },
	"stderr": func() Output { return &Stderr{} },
	"file":   func() Output { return &File{} },
}

type Output interface {
	Write([]byte) error
	GetSerializers() []string
}

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

type File struct {
	Serializers []string `toml:"serializers"`
	Filename    string   `toml:"Filename"`
}

func (f *File) GetSerializers() []string {
	return f.Serializers
}

func (f *File) Write(p []byte) error {
	return os.WriteFile(f.Filename, p, 0644)
}
