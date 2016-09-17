package goo

import (
	"bytes"
	"go/format"
	"regexp"
	"text/template"
)

func Macro(name string, file []byte, data interface{}) ([]byte, error) {
	file = regexp.MustCompile(`(?m:^\/\/go:generate.+$)`).ReplaceAll(file, []byte(""))
	file = regexp.MustCompile(`__(\w[\w_]*)__`).ReplaceAll(file, []byte("{{.$1}}"))
	file = regexp.MustCompile(`\/\/\/(.*)`).ReplaceAll(file, []byte("$1"))
	file = regexp.MustCompile(`\/\*\*(.*)\*\*\/`).ReplaceAll(file, []byte("$1"))

	var t, err = template.New(name).Parse(string(file))

	if err != nil {
		return nil, err
	}

	var b = &bytes.Buffer{}

	if err := t.Execute(b, data); err != nil {
		return nil, err
	}

	if file, err = format.Source(b.Bytes()); err != nil {
		return nil, err
	}

	return file, nil
}
