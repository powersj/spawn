package serializers

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/mitchellh/mapstructure"
)

var Registry = map[string]func() Serializer{
	"template": func() Serializer { return &Template{} },
}

type Serializer interface {
	Serialize(template.FuncMap) []byte
}

func Load(conf map[string][]map[string]interface{}) (map[string]Serializer, error) {
	serials := make(map[string]Serializer)
	for serialType, serialConfigs := range conf {
		constructor, exists := Registry[serialType]
		if !exists {
			return nil, fmt.Errorf("Serializer type %s not found", serialType)
		}
		for _, serialConf := range serialConfigs {
			serial := constructor()
			if err := mapstructure.Decode(serialConf, serial); err != nil {
				return nil, fmt.Errorf("Error decoding config: %w", err)
			}
			if id, ok := serialConf["id"].(string); ok {
				serials[id] = serial
			} else {
				return nil, fmt.Errorf("ID missing for serializer %v", serialConf)
			}
		}
	}

	return serials, nil
}

type Template struct {
	Template string `toml:"template"`
}

func (t *Template) Serialize(funcMap template.FuncMap) []byte {
	tmpl, err := template.New("example").Funcs(funcMap).Parse(t.Template)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, nil); err != nil {
		panic(err)
	}

	return out.Bytes()
}
