package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

type agentConfig struct {
	Debug     bool          `toml:"debug"`
	Interval  time.Duration `toml:"interval"`
	PprofPort string        `toml:"pprof_port"`
	Seed      int64         `toml:"seed"`
	Trace     bool          `toml:"trace"`
}

type Config struct {
	Agent      agentConfig                 `toml:"agent"`
	Generator  map[string][]map[string]any `toml:"generator"`
	Serializer map[string][]map[string]any `toml:"serializer"`
	Output     map[string][]map[string]any `toml:"output"`
}

func NewConfig() (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("examples/agent.toml", &conf); err != nil {
		return nil, err
	}

	if conf.Agent.Interval == 0 {
		conf.Agent.Interval = 5 * time.Second
	}

	return &conf, nil
}
