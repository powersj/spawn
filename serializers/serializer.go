package serializers

import (
	"bytes"
	"text/template"
)

var Registry = map[string]func() Serializer{
	"template": func() Serializer { return &Template{} },
}

type Serializer interface {
	Serialize(template.FuncMap) []byte
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
