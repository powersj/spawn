package generators

import (
	"fmt"
	"text/template"

	"github.com/mitchellh/mapstructure"
)

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
