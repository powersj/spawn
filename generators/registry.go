package generators

import "github.com/powersj/spawn/generators/randfloat64"

var Registry = map[string]func() Generator{
	"randomfloat64": func() Generator { return &randfloat64.Randomfloat64{} },
	"randomint64":   func() Generator { return &randfloat64.Randomfloat64{} },
}
