package serializers

import "github.com/powersj/spawn/serializers/template"

var Registry = map[string]func() Serializer{
	"template": func() Serializer { return &template.Template{} },
}
