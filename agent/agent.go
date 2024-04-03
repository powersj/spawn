package agent

import (
	"fmt"
	"os"
	"slices"
	"text/template"
	"time"

	"github.com/powersj/spawn/config"
	"github.com/powersj/spawn/generators"
	"github.com/powersj/spawn/outputs"
	"github.com/powersj/spawn/outputs/stdout"
	"github.com/powersj/spawn/serializers"
)

func NewAgent(conf *config.Config) (*Agent, error) {
	return &Agent{
		config: *conf,
	}, nil
}

func NewPprofServer() *PprofServer {
	return &PprofServer{
		err: make(chan error),
	}
}

type Agent struct {
	config config.Config
}

func (a *Agent) Run() error {
	funcMap, err := generators.Load(a.config.Generator)
	if err != nil {
		return fmt.Errorf("Error loading generators: %w", err)
	}

	outs, err := outputs.Load(a.config.Output)
	if err != nil {
		return fmt.Errorf("Error loading outputs: %w", err)
	}

	serials, err := serializers.Load(a.config.Serializer)
	if err != nil {
		return fmt.Errorf("Error loading serializers: %w", err)

	}

	ticker := time.NewTicker(a.config.Agent.Interval)
	for {
		<-ticker.C
		for serialID, serial := range serials {
			go a.generate(serialID, serial, funcMap, outs)
		}
	}
}

func (a *Agent) RunOnce() error {
	o := stdout.Stdout{}

	funcMap, err := generators.Load(a.config.Generator)
	if err != nil {
		return fmt.Errorf("Error loading generators: %w", err)
	}

	serials, err := serializers.Load(a.config.Serializer)
	if err != nil {
		return fmt.Errorf("Error loading serializers: %w", err)

	}

	for _, serial := range serials {
		out := serial.Serialize(funcMap)
		if err := o.Write(out); err != nil {
			return fmt.Errorf("Error writing output: %w", err)
		}
	}

	return nil
}

func (a *Agent) generate(
	serialID string,
	serial serializers.Serializer,
	funcMap template.FuncMap,
	outs map[string]outputs.Output,
) {
	out := serial.Serialize(funcMap)
	for _, o := range outs {
		if slices.Contains(o.GetSerializers(), serialID) {
			if err := o.Write(out); err != nil {
				fmt.Println("Error writing output:", err)
				os.Exit(1)
			}
		}
	}
}
