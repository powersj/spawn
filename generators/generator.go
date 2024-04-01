package generators

import (
	"fmt"
	"math/rand"
	"text/template"

	"github.com/mitchellh/mapstructure"
)

var Registry = map[string]func() Generator{
	"randomfloat64": func() Generator { return &Randomfloat64{} },
	"randomint64":   func() Generator { return &Randomint64{} },
}

type Generator interface {
	Generate() any
}

func Load(conf map[string][]map[string]any) (template.FuncMap, error) {
	funcMap := make(template.FuncMap)

	for genType, genConfigs := range conf {
		constructor, exists := Registry[genType]
		if !exists {
			fmt.Printf("Generator type %s not found\n", genType)
			continue
		}
		for _, genConf := range genConfigs {
			generator := constructor()
			if err := mapstructure.Decode(genConf, generator); err != nil {
				fmt.Println("Error decoding config:", err)
				continue
			}
			if id, ok := genConf["id"].(string); ok {
				funcMap[id] = generator.Generate
			} else {
				fmt.Println("ID missing for generator", genConf)
			}
		}
	}

	return funcMap, nil
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
