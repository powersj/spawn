package serializers

import (
	"fmt"
	txtTemplate "text/template"

	"github.com/mitchellh/mapstructure"
)

type Serializer interface {
	Serialize(txtTemplate.FuncMap) []byte
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
