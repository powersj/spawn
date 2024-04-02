package template

import (
	"bytes"
	txtTemplate "text/template"
)

type Template struct {
	Template string `toml:"template"`
}

func (t *Template) Serialize(funcMap txtTemplate.FuncMap) []byte {
	tmpl, err := txtTemplate.New("example").Funcs(funcMap).Parse(t.Template)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, nil); err != nil {
		panic(err)
	}

	return out.Bytes()
}
