package outputs

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

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
