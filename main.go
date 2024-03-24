package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Agent      Agent      `toml:"agent"`
	Generator  Generator  `toml:"generator"`
	Serializer Serializer `toml:"serializer"`
	Output     Output     `toml:"output"`
}

func main() {
	configPath := "examples/agent.toml"
	f, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var config Config
	decoder := toml.NewDecoder(bytes.NewReader(f)).DisallowUnknownFields()
	if err := decoder.Decode(&config); err != nil {
		var strict *toml.StrictMissingError
		if errors.As(err, &strict) {
			panic(fmt.Errorf("unrecognized config option found in %q:\n%s", configPath, strict.String()))
		}
		panic(fmt.Errorf("unable to unmarshal config file %q: %w", configPath, err))
	}

	fmt.Printf("configuration: %v\n", config)

	for _, serializer := range config.Serializer.Templates {
		funcMap := template.FuncMap{
			"float": config.Generator.RandomFloat64[0].Generate,
			//"int":   config.Generator.RandomInt64[0].Generate,
		}

		tmpl, err := template.New("example").Funcs(funcMap).Parse(serializer.Template)
		if err != nil {
			panic(err)
		}

		var out bytes.Buffer
		if err := tmpl.Execute(&out, nil); err != nil {
			panic(err)
		}

		//os.Stdout.Write(out.Bytes())
		fmt.Println(out.String())
	}
}
