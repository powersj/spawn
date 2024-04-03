package outputs

import (
	"github.com/powersj/spawn/outputs/file"
	"github.com/powersj/spawn/outputs/stderr"
	"github.com/powersj/spawn/outputs/stdout"
)

var Registry = map[string]func() Output{
	"stdout": func() Output { return &stdout.Stdout{} },
	"stderr": func() Output { return &stderr.Stderr{} },
	"file":   func() Output { return &file.File{} },
}
