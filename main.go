package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
	"github.com/powersj/spawn/generators"
	"github.com/powersj/spawn/serializers"
)

type Config struct {
	Agent      Agent                       `toml:"agent"`
	Generator  map[string][]map[string]any `toml:"generator"`
	Serializer map[string][]map[string]any `toml:"serializer"`
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("examples/agent.toml", &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gens := make(map[string]generators.Generator)
	funcMap := make(template.FuncMap)
	for genType, genConfigs := range conf.Generator {
		constructor, exists := generators.Registry[genType]
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
				gens[id] = generator
				funcMap[id] = generator.Generate
			} else {
				fmt.Println("ID missing for generator", genConf)
			}
		}
	}

	serials := make(map[string]serializers.Serializer)
	for serialType, serialConfig := range conf.Serializer {
		constructor, exists := serializers.Registry[serialType]
		if !exists {
			fmt.Printf("Generator type %s not found\n", serialType)
			continue
		}
		for _, genConf := range serialConfig {
			generator := constructor()
			if err := mapstructure.Decode(genConf, generator); err != nil {
				fmt.Println("Error decoding config:", err)
				continue
			}
			if id, ok := genConf["id"].(string); ok {
				serials[id] = generator
			} else {
				fmt.Println("ID missing for generator", genConf)
			}
		}

		for _, serial := range serials {
			out := serial.Serialize(funcMap)
			fmt.Println(out)
		}

	}
}
