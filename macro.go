package goo

import (
	"bytes"
	"go/format"
	"regexp"
	"text/template"
)

var (
	genRegexp       = regexp.MustCompile(`(?m:^\/\/go:generate.+$)`)
	gooRegexp       = regexp.MustCompile(`__(?i:goo)_(?i:comment)_[\w\d]+__`)
	identRegexp     = regexp.MustCompile(`__([[:alpha:]](_?[[:alnum:]])*)__`)
	lineRegexp      = regexp.MustCompile(`\/\/\/(.*)`)
	multilineRegexp = regexp.MustCompile(`\/\*\*(.*)\*\*\/`)
)

func preprocessIdents(bs []byte) []byte {
	var b bytes.Buffer

	b.WriteString("{{.")
	b.Write(bytes.Replace(bs[2:len(bs)-2], []byte("_"), []byte("."), -1))
	b.WriteString("}}")

	return b.Bytes()
}

// Macro calls MacroPreprocess, MacroProcess, and then MacroFormat with file and
// returns the result.
func Macro(name string, file []byte, data interface{}) ([]byte, error) {
	file, err := MacroProcess(name, MacroPreprocess(file), data)

	if err != nil {
		return nil, err
	}

	return MacroFormat(file)
}

// MacroFormat returns file formatted with go/format.
func MacroFormat(file []byte) ([]byte, error) {
	return format.Source(file)
}

// MacroPreprocess returns file with generate comments stripped, identifiers of
// the form __My_ID__ renamed to {{.My_ID}}, and single-line comments beginning
// with a slash and multi-line comments beginning and ending with an asterisk
// replaced with their content.
func MacroPreprocess(file []byte) []byte {
	file = genRegexp.ReplaceAll(file, []byte(""))
	file = gooRegexp.ReplaceAll(file, []byte(""))
	file = identRegexp.ReplaceAllFunc(file, preprocessIdents)
	file = lineRegexp.ReplaceAll(file, []byte("$1"))
	file = multilineRegexp.ReplaceAll(file, []byte("$1"))

	return file
}

// MacroProcess returns file ran as a text/template with name and data.
func MacroProcess(name string, file []byte, data interface{}) ([]byte, error) {
	var t, err = template.New(name).Parse(string(file))

	if err != nil {
		return nil, err
	}

	var b = &bytes.Buffer{}

	if err := t.Execute(b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
