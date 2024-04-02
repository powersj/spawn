package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/powersj/spawn/agent"
	"github.com/powersj/spawn/config"
	"github.com/powersj/spawn/internal"
	"github.com/urfave/cli/v2"
)

const (
	cliDescription = `spawn is a TOML config driven agent to generate data

The configuration consists of generators, serializers, and outputs. Generators
are used to generate random data. Serializers determine a specific output data
type and format, and outputs is where the data is sent.

Agent level settings are also available to control logging and generation of
data.
`
)

func main() {
	app := &cli.App{
		Name:        "spawn",
		Usage:       "TOML config driven agent to generate data",
		Description: cliDescription,
		Suggest:     true,
		Commands: []*cli.Command{
			{
				Name:        "run",
				Usage:       "Runs the generator",
				Description: "Runs the generator",
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() == 0 {
						return fmt.Errorf("missing TOML file")
					} else if ctx.NArg() > 1 {
						return fmt.Errorf("too many arguments")
					}

					c, _ := config.NewConfig()
					a, _ := agent.NewAgent(c)

					pprof := agent.NewPprofServer()
					pprof.Start(c.Agent.PprofPort)

					return a.Run()
				},
			},
			{
				Name:        "toml",
				Usage:       "Used to verify a TOML configuration format",
				Description: "Used to verify a TOML configuration format",
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() == 0 {
						return fmt.Errorf("missing TOML file")
					} else if ctx.NArg() > 1 {
						return fmt.Errorf("too many arguments")
					}

					var data any
					if _, err := toml.DecodeFile(ctx.Args().First(), &data); err != nil {
						return fmt.Errorf("Invalid TOML File: %w", err)
					}

					fmt.Println("The TOML file is valid.")
					return nil
				},
			},
			{
				Name:        "once",
				Usage:       "Run all generators and serializers and output to stdout",
				Description: "Run all generators and serializers and output to stdout",
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() == 0 {
						return fmt.Errorf("missing TOML file")
					} else if ctx.NArg() > 1 {
						return fmt.Errorf("too many arguments")
					}

					c, _ := config.NewConfig()
					a, _ := agent.NewAgent(c)

					fmt.Println("skipping outputs in run once mode")
					return a.RunOnce()
				},
			},
			{
				Name:        "version",
				Usage:       "Print version, build, and platform info",
				Description: "Print version, build, and platform info",
				Action: func(*cli.Context) error {
					fmt.Println(internal.AppVersion())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
