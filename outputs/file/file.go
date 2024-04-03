package file

import "os"

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
