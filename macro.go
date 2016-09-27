package goo

import (
	"bytes"
	"go/format"
	"regexp"
	"text/template"
)

var (
	macroEscape     = regexp.MustCompile(`ESCAPE_([[:alnum:]]+)`)
	macroEscapeX    = regexp.MustCompile(`ESCAPEX_(\w+)_END`)
	macroEscaped    = regexp.MustCompile(`\{\{\/\*ESCAPE\*\/\}\}`)
	macroEscapedX   = regexp.MustCompile(`\{\{\/\*ESCAPEX\*\/\}\}`)
	macroGenerate   = regexp.MustCompile(`(?m:^\/\/go:generate.*$)`)
	macroIdentifier = regexp.MustCompile(`__[A-Za-z](_?[[:alnum:]])*__`)
	macroLine       = regexp.MustCompile(`\/\/\/(.*)`)
	macroMultiLine  = regexp.MustCompile(`(?s:\/\*\*(.*)\*\*\/)`)
	macroShorthand  = regexp.MustCompile(`(?m:^__([[:alnum:]]+)__$)`)
)

var (
	macroDollarOne       = []byte("$1")
	macroEmpty           = []byte("")
	macroPlaceholder     = []byte(`{{/*ESCAPE*/}}`)
	macroPlaceholderX    = []byte(`{{/*ESCAPEX*/}}`)
	macroShorthandResult = []byte(`{{.$1}}`)
	macroSpace           = []byte(" ")
	macroUnderscore      = []byte("_")
	macroUnderscores     = []byte("__")
)

var macroAliases = map[string][]byte{
	"ASSIGN":   []byte(":="),
	"DOT":      []byte("{{.}}"),
	"ELSE":     []byte("{{else}}"),
	"END":      []byte("{{end}}"),
	"LBRACES":  []byte("{{"),
	"LCOMMENT": []byte("{{/*"),
	"RBRACES":  []byte("}}"),
	"RCOMMENT": []byte("*/}}"),
	"TRIM":     []byte("-"),
}

var macroDirs = map[*regexp.Regexp][]byte{
	regexp.MustCompile(`(?m:^__ACTION_(\w+)__$)`):          []byte("{{$1}}"),
	regexp.MustCompile(`(?m:^__CUSTOM_(\w+)__$)`):          []byte("$1"),
	regexp.MustCompile(`(?m:^__FIELD_([[:alnum:]]+)__$)`):  []byte("{{.$1}}"),
	regexp.MustCompile(`(?m:^__KEY_([[:alnum:]]+)__$)`):    []byte("{{.$1}}"),
	regexp.MustCompile(`(?m:^__METHOD_([[:alnum:]]+)__$)`): []byte("{{.$1}}"),
	regexp.MustCompile(`(?m:^__VAR_([[:alnum:]]+)__$)`):    []byte("{{$$$1}}"),
}

var macroFuncs = map[*regexp.Regexp][]byte{
	regexp.MustCompile(`ACTIONX_(\w+)_END`):                                   []byte("{{$1}}"),
	regexp.MustCompile(`ACTION_([[:alnum:]]+)`):                               []byte("{{$1}}"),
	regexp.MustCompile(`BLOCKX_(\w+)_BEGIN_(\w+)_END`):                        []byte(`{{block "$1" $2}}`),
	regexp.MustCompile(`BLOCK_([[:alnum:]]+)_([[:alnum:]]+)`):                 []byte(`{{block "$1" $2}}`),
	regexp.MustCompile(`FIELD_([[:alnum:]]+)`):                                []byte(".$1"),
	regexp.MustCompile(`GROUPX_(\w+)_END`):                                    []byte("($1)"),
	regexp.MustCompile(`GROUP_([[:alnum:]]+)`):                                []byte("($1)"),
	regexp.MustCompile(`IFEX_(\w+)_THEN_(\w+)_ELSE_(\w+)_END`):                []byte("{{if $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile(`IFE_([[:alnum:]]+)_([[:alnum:]]+)_([[:alnum:]]+)`):    []byte("{{if $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile(`IFX_(\w+)_THEN_(\w+)_END`):                            []byte("{{if $1}}$2{{end}}"),
	regexp.MustCompile(`IF_([[:alnum:]]+)_([[:alnum:]]+)`):                    []byte("{{if $1}}$2{{end}}"),
	regexp.MustCompile(`INITX_([[:alnum:]]+)_(\w+)_END`):                      []byte("$$$1 := $2"),
	regexp.MustCompile(`INIT_([[:alnum:]]+)_([[:alnum:]]+)`):                  []byte("$$$1 := .$2"),
	regexp.MustCompile(`KEY_([[:alnum:]]+)`):                                  []byte(".$1"),
	regexp.MustCompile(`METHOD_([[:alnum:]]+)`):                               []byte(".$1"),
	regexp.MustCompile(`OMITX_\w+_END`):                                       []byte(""),
	regexp.MustCompile(`OMIT_[[:alnum:]]+`):                                   []byte(""),
	regexp.MustCompile(`RANGEEX_(\w+)_BEGIN_(\w+)_ELSE_(\w+)_END`):            []byte("{{range $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile(`RANGEE_([[:alnum:]]+)_([[:alnum:]]+)_([[:alnum:]]+)`): []byte("{{range $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile(`RANGEX_(\w+)_BEGIN_(\w+)_END`):                        []byte("{{range $1}}$2{{end}}"),
	regexp.MustCompile(`RANGE_([[:alnum:]]+)_([[:alnum:]]+)`):                 []byte("{{range $1}}$2{{end}}"),
	regexp.MustCompile(`RAWX_(\w+)_END`):                                      []byte("`$1`"),
	regexp.MustCompile(`RAW_([[:alnum:]]+)`):                                  []byte("`$1`"),
	regexp.MustCompile(`RUNEO_([0-7]{3})`):                                    []byte(`\$1`),
	regexp.MustCompile(`RUNEU4_([0-9A-Fa-f]{4})`):                             []byte(`\u$1`),
	regexp.MustCompile(`RUNEU8_([0-9A-Fa-f]{8})`):                             []byte(`\U$1`),
	regexp.MustCompile(`RUNEX_([0-9A-Fa-f]{2})`):                              []byte(`\x$1`),
	regexp.MustCompile(`RUNE_([[:alnum:]]+)`):                                 []byte("'$1'"),
	regexp.MustCompile(`STRINGX_(\w+)_END`):                                   []byte(`"$1"`),
	regexp.MustCompile(`STRING_([[:alnum:]]+)`):                               []byte(`"$1"`),
	regexp.MustCompile(`TEMPLATEX_(\w+)_BEGIN_(\w+)_END`):                     []byte(`{{template "$1" $2}}`),
	regexp.MustCompile(`TEMPLATE_([[:alnum:]]+)_([[:alnum:]]+)`):              []byte(`{{template "$1" $2}}`),
	regexp.MustCompile(`UNICODE_([[:alnum:]]+)`):                              []byte(`'\u$1'`),
	regexp.MustCompile(`VAR_([[:alnum:]]+)`):                                  []byte("$$$1"),
	regexp.MustCompile(`WITHEX_(\w+)_BEGIN_(\w+)_ELSE_(\w+)_END`):             []byte("{{with $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile(`WITHE_([[:alnum:]]+)_([[:alnum:]]+)_([[:alnum:]]+)`):  []byte("{{with $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile(`WITHX_(\w+)_BEGIN_(\w+)_END`):                         []byte("{{with $1}}$2{{end}}"),
	regexp.MustCompile(`WITH_([[:alnum:]]+)_([[:alnum:]]+)`):                  []byte("{{with $1}}$2{{end}}"),
}

var macroSyms = map[string][]byte{
	"AMPERSAND":   []byte("&"),
	"APOSTROPHE":  []byte("'"),
	"ASTERISK":    []byte("*"),
	"AT":          []byte("@"),
	"BACKSLASH":   []byte(`\`),
	"CARET":       []byte("^"),
	"COLON":       []byte(":"),
	"COMMA":       []byte(","),
	"DOLLAR":      []byte("$"),
	"EQUALS":      []byte("="),
	"EXCLAMATION": []byte("!"),
	"GRAVE":       []byte("`"),
	"GREATER":     []byte(">"),
	"LBRACE":      []byte("{"),
	"LBRACKET":    []byte("["),
	"LESS":        []byte("<"),
	"LPAREN":      []byte("("),
	"MINUS":       []byte("-"),
	"NUMBER":      []byte("#"),
	"PERCENT":     []byte("%"),
	"PERIOD":      []byte("."),
	"PIPE":        []byte("|"),
	"PLUS":        []byte("+"),
	"QUESTION":    []byte("?"),
	"QUOTATION":   []byte(`"`),
	"RBRACE":      []byte("}"),
	"RBRACKET":    []byte("]"),
	"RPAREN":      []byte(")"),
	"SEMICOLON":   []byte(";"),
	"SLASH":       []byte("/"),
	"SPACE":       []byte(" "),
	"TILDE":       []byte("~"),
	"UNDERSCORE":  []byte("_"),
}

func macroEscapeFunc(bs []byte) []byte {
	return macroPlaceholder
}

func macroEscapeXFunc(bs []byte) []byte {
	return macroPlaceholderX
}

func macroPreprocess(bs []byte) []byte {
	var escapes = macroEscape.FindAll(bs, -1)
	var escapesX = macroEscapeX.FindAll(bs, -1)

	bs = macroEscape.ReplaceAllFunc(bs, macroEscapeFunc)
	bs = macroEscapeX.ReplaceAllFunc(bs, macroEscapeXFunc)

	var dir bool

	for prefix, result := range macroDirs {
		if prefix.Match(bs) {
			bs = prefix.ReplaceAll(bs, result)
			dir = true

			break
		}
	}

	if !dir {
		bs = macroShorthand.ReplaceAll(bs, macroShorthandResult)
	}

	for call, result := range macroFuncs {
		bs = call.ReplaceAll(bs, result)
	}

	for short, long := range macroAliases {
		bs = bytes.Replace(bs, []byte(short), long, -1)
	}

	bs = bytes.TrimPrefix(bs, macroUnderscores)
	bs = bytes.TrimSuffix(bs, macroUnderscores)
	bs = bytes.Replace(bs, macroUnderscore, macroSpace, -1)

	for name, symbol := range macroSyms {
		bs = bytes.Replace(bs, []byte(name), symbol, -1)
	}

	bs = macroEscaped.ReplaceAllFunc(bs, func([]byte) []byte {
		var bs = escapes[0]

		escapes = escapes[1:]

		return macroEscape.ReplaceAll(bs, macroDollarOne)
	})

	bs = macroEscapedX.ReplaceAllFunc(bs, func([]byte) []byte {
		var bs = escapesX[0]

		escapesX = escapesX[1:]

		return macroEscapeX.ReplaceAll(bs, macroDollarOne)
	})

	return bs
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
	file = macroGenerate.ReplaceAll(file, macroEmpty)
	file = macroIdentifier.ReplaceAllFunc(file, macroPreprocess)
	file = macroLine.ReplaceAll(file, macroDollarOne)
	file = macroMultiLine.ReplaceAll(file, macroDollarOne)

	return file
}

// MacroProcess returns file ran as a text/template with name and data.
func MacroProcess(name string, file []byte, data interface{}) ([]byte, error) {
	var t, err = template.New(name).Parse(string(file))

	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
