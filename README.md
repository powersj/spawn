# Spawn

*TOML config driven agent to generate data*

[![CircleCI](https://circleci.com/gh/powersj/spawn.svg?style=svg)](https://circleci.com/gh/powersj/spawn) [![Go Reference](https://pkg.go.dev/badge/github.com/powersj/spawn.svg)](https://pkg.go.dev/github.com/powersj/spawn)

## Overview

This is an agent to generate data via a TOML configuration file. Users define
a serializer format and call generator functions to generate data at each
interval. The generated data is then sent the corresponding outputs.

To create a basic JSON with some random numeric values:

```TOML
[[generator.randomfloat64]]
    id = "float" # id used to reference in serializer template
[[serializer.template]]
    id = "json"  # id used to reference in outputs
    template = """{ "value": {{ float }} }"""
[[output.stderr]]
    serializers = ["json"]
```

Would produce a different JSON value at each interval in the format:

```json
{ "value": 123.456 }
```

## Install

Pre-build binaries are available on the releases page.

## Usage

To generate data use the `run` subcommand:

```sh
./spawn run <toml file>
```

The other subcommands include:

* `once`: To run all serializers and generators once and output to stdout
* `toml`: To validate a TOML file

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.

Please make sure to update tests as appropriate.

## Support

If you find a bug, have a question, or ideas for improvements please file an
[issue](https://github.com/powersj/spawn/issues/new) on GitHub.
