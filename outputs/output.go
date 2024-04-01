package outputs

import (
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
)

var Registry = map[string]func() Output{
	"stdout": func() Output { return &Stdout{} },
	"stderr": func() Output { return &Stderr{} },
	"file":   func() Output { return &File{} },
}

type Output interface {
	Write([]byte) error
	GetSerializers() []string
}

func Load(conf map[string][]map[string]interface{}) (map[string]Output, error) {
	outs := make(map[string]Output)
	for outputType, outputConfig := range conf {
		constructor, exists := Registry[outputType]
		if !exists {
			fmt.Printf("Output type %s not found\n", outputType)
			continue
		}
		for _, outConf := range outputConfig {
			output := constructor()
			if err := mapstructure.Decode(outConf, output); err != nil {
				fmt.Println("Error decoding config:", err)
				continue
			}
			outs[outputType] = output
		}
	}

	return outs, nil
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
