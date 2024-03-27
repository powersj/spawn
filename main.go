package main

import (
	"fmt"
	"html/template"
	"os"
	"slices"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
	"github.com/powersj/spawn/generators"
	"github.com/powersj/spawn/outputs"
	"github.com/powersj/spawn/serializers"
)

type Config struct {
	Agent      Agent                       `toml:"agent"`
	Generator  map[string][]map[string]any `toml:"generator"`
	Serializer map[string][]map[string]any `toml:"serializer"`
	Output     map[string][]map[string]any `toml:"output"`
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("examples/agent.toml", &conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
				funcMap[id] = generator.Generate
			} else {
				fmt.Println("ID missing for generator", genConf)
			}
		}
	}

	outs := make(map[string]outputs.Output)
	for outputType, outputConfig := range conf.Output {
		constructor, exists := outputs.Registry[outputType]
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
	}

	ticker := time.NewTicker(time.Duration(5 * time.Second))
	for {
		<-ticker.C
		for serial_id, serial := range serials {
			go generate(serial_id, serial, funcMap, outs)
		}
	}
}

func generate(serial_id string, serial serializers.Serializer, funcMap template.FuncMap, outs map[string]outputs.Output) {
	out := serial.Serialize(funcMap)
	for _, o := range outs {
		if slices.Contains(o.GetSerializers(), serial_id) {
			if err := o.Write(out); err != nil {
				fmt.Println("Error writing output:", err)
				os.Exit(1)
			}
		}
	}
}
