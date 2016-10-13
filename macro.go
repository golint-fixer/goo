package goo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"regexp"
	"text/template"
)

const macroIdentifierContent = "[[:alnum:]](?:_?[[:alnum:]])*?"

var (
	macroEscape      = regexp.MustCompile("ESCAPE_([[:alnum:]]+)")
	macroEscapeX     = regexp.MustCompile("ESCAPEX_(" + macroIdentifierContent + ")_ENDESCAPE")
	macroEscaped     = regexp.MustCompile(`\{\{\/\*GOO:ESCAPE\*\/\}\}`)
	macroEscapedX    = regexp.MustCompile(`\{\{\/\*GOO:ESCAPEX\*\/\}\}`)
	macroFields      = regexp.MustCompile("FIELDS_(" + macroIdentifierContent + ")_ENDFIELDS")
	macroGenerate    = regexp.MustCompile(`(?m:^\/\/go:generate.*$)`)
	macroIdentifier  = regexp.MustCompile("__(" + macroIdentifierContent + ")__")
	macroKeys        = regexp.MustCompile("KEYS_(" + macroIdentifierContent + ")_ENDKEYS")
	macroLine        = regexp.MustCompile(`\/\/\/(.*)`)
	macroLonghand    = regexp.MustCompile("(?m:^__X_(" + macroIdentifierContent + ")__$)")
	macroMethods     = regexp.MustCompile("METHODS_(" + macroIdentifierContent + ")_ENDMETHODS")
	macroMultiLine   = regexp.MustCompile(`(?s:\/\*\*(.*?)\*\*\/)`)
	macroShorthand   = regexp.MustCompile("(?m:^__(" + macroIdentifierContent + ")__$)")
	macroUnderscoreR = regexp.MustCompile("_")
)

var (
	macroDollarOne       = []byte("$1")
	macroDollarOneBraces = []byte("{{$1}}")
	macroDot             = []byte(".")
	macroEmpty           = []byte("")
	macroPlaceholder     = []byte("{{/*GOO:ESCAPE*/}}")
	macroPlaceholderX    = []byte("{{/*GOO:ESCAPEX*/}}")
	macroSpace           = []byte(" ")
	macroUnderscoreB     = []byte("_")
	macroUnderscores     = []byte("__")
)

var macroAliases = map[string][]byte{
	"ASSIGN":   []byte(":="),
	"DOT":      []byte("{{.}}"),
	"EMPTY":    []byte(""),
	"LBRACES":  []byte("{{"),
	"LCOMMENT": []byte("{{/*"),
	"RBRACES":  []byte("}}"),
	"RCOMMENT": []byte("*/}}"),
	"TRIM":     []byte("-"),
}

var macroFuncs = map[*regexp.Regexp][]byte{
	regexp.MustCompile("ACTION_(.+?)_ENDACTION"):                       []byte("{{$1}}"),
	regexp.MustCompile("BLOCK_(.+?)_BEGIN_(.+?)_ENDBLOCK"):             []byte(`{{block "$1" $2}}`),
	regexp.MustCompile("FIELDX_(.+?)_ENDFIELD"):                        []byte(".$1"),
	regexp.MustCompile("FIELD_([[:alnum:]]+)"):                         []byte(".$1"),
	regexp.MustCompile("GROUPX_(.+?)_ENDGROUP"):                        []byte("($1)"),
	regexp.MustCompile("GROUP_([[:alnum:]]+)"):                         []byte("($1)"),
	regexp.MustCompile("IFE_(.+?)_THEN_(.+?)_ELSE_(.+?)_ENDIF"):        []byte("{{if $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile("IF_(.+?)_THEN_(.+?)_ENDIF"):                    []byte("{{if $1}}$2{{end}}"),
	regexp.MustCompile("INIT_([[:alnum:]]+)_(.+?)_ENDINIT"):            []byte("$$$1 := $2"),
	regexp.MustCompile("KEYX_(.+?)_ENDKEY"):                            []byte(".$1"),
	regexp.MustCompile("KEY_([[:alnum:]]+)"):                           []byte(".$1"),
	regexp.MustCompile("METHODX_(.+?)_ENDMETHOD"):                      []byte(".$1"),
	regexp.MustCompile("METHOD_([[:alnum:]]+)"):                        []byte(".$1"),
	regexp.MustCompile("OMITX_.+_ENDOMIT"):                             []byte(""),
	regexp.MustCompile("OMIT_[[:alnum:]]+"):                            []byte(""),
	regexp.MustCompile("RANGEE_(.+?)_BEGIN_(.+?)_ELSE_(.+?)_ENDRANGE"): []byte("{{range $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile("RANGE_(.+?)_BEGIN_(.+?)_ENDRANGE"):             []byte("{{range $1}}$2{{end}}"),
	regexp.MustCompile("RAWX_(.+?)_ENDRAW"):                            []byte("`$1`"),
	regexp.MustCompile("RAW_([[:alnum:]]+)"):                           []byte("`$1`"),
	regexp.MustCompile("RUNEO_([0-7]{3})"):                             []byte(`\$1`),
	regexp.MustCompile("RUNEU4_([0-9A-Fa-f]{4})"):                      []byte(`\u$1`),
	regexp.MustCompile("RUNEU8_([0-9A-Fa-f]{8})"):                      []byte(`\U$1`),
	regexp.MustCompile("RUNEX_([0-9A-Fa-f]{2})"):                       []byte(`\x$1`),
	regexp.MustCompile("RUNE_([[:alnum:]]+)"):                          []byte("'$1'"),
	regexp.MustCompile("STRINGX_(.+?)_ENDSTRING"):                      []byte(`"$1"`),
	regexp.MustCompile("STRING_([[:alnum:]]+)"):                        []byte(`"$1"`),
	regexp.MustCompile("TEMPLATE_(.+?)_BEGIN_(.+?)_ENDTEMPLATE"):       []byte(`{{template "$1" $2}}`),
	regexp.MustCompile("UNICODE_([[:alnum:]]+)"):                       []byte(`'\u$1'`),
	regexp.MustCompile("VAR_([[:alnum:]]+)"):                           []byte("$$$1"),
	regexp.MustCompile("WITHE_(.+?)_BEGIN_(.+?)_ELSE_(.+?)_ENDWITH"):   []byte("{{with $1}}$2{{else}}$3{{end}}"),
	regexp.MustCompile("WITH_(.+?)_BEGIN_(.+?)_ENDWITH"):               []byte("{{with $1}}$2{{end}}"),
}

var macroSpecials = map[*regexp.Regexp]func([]byte) []byte{
	macroFields:  macroFieldsFunc,
	macroKeys:    macroKeysFunc,
	macroMethods: macroMethodsFunc,
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

func macroFieldsFunc(bs []byte) []byte {
	return append([]byte("."), macroUnderscoreR.ReplaceAll(macroFields.FindSubmatch(bs)[1], macroDot)...)
}

func macroKeysFunc(bs []byte) []byte {
	return append([]byte("."), macroUnderscoreR.ReplaceAll(macroKeys.FindSubmatch(bs)[1], macroDot)...)
}

func macroMethodsFunc(bs []byte) []byte {
	return append([]byte("."), macroUnderscoreR.ReplaceAll(macroMethods.FindSubmatch(bs)[1], macroDot)...)
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
	bs = macroLonghand.ReplaceAll(bs, macroDollarOne)
	bs = macroShorthand.ReplaceAll(bs, macroDollarOneBraces)

	for call, f := range macroSpecials {
		bs = call.ReplaceAllFunc(bs, f)
	}

	for call, result := range macroFuncs {
		bs = call.ReplaceAll(bs, result)
	}

	for short, long := range macroAliases {
		bs = bytes.Replace(bs, []byte(short), long, -1)
	}

	bs = bytes.TrimPrefix(bs, macroUnderscores)
	bs = bytes.TrimSuffix(bs, macroUnderscores)
	bs = bytes.Replace(bs, macroUnderscoreB, macroSpace, -1)

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

func macroCommentGroup(cg *ast.CommentGroup) []string {
	if cg == nil {
		return nil
	}

	var cs []string

	for _, d := range cg.List {
		cs = append(cs, d.Text)
	}

	return cs
}

type MacroInterface struct {
	Doc       []string
	Methods   []*MacroMethod
	Name      string
	Package   string
	Qualifier string
}

func GetMacroInterface(package_, identifier string) (*MacroInterface, error) {
	var bp, err = build.Import(package_, "", 0)

	if err != nil {
		return nil, fmt.Errorf("cannot find package %v: %v", package_, err)
	}

	aps, err := parser.ParseDir(token.NewFileSet(), bp.Dir, nil, parser.ParseComments)

	if err != nil {
		return nil, fmt.Errorf("cannot parse package %v: %v", package_, err)
	}

	var match *ast.TypeSpec

	for _, ap := range aps {
		ast.Inspect(ap, func(n ast.Node) bool {
			if n == nil || match != nil {
				return false
			}

			var g, ok = n.(*ast.GenDecl)

			if !ok {
				return true
			}

			if g.Tok != token.TYPE {
				return false
			}

			for _, s := range g.Specs {
				var t, ok = s.(*ast.TypeSpec)

				if !ok {
					continue
				}

				if _, ok := t.Type.(*ast.InterfaceType); !ok {
					continue
				}

				if t.Name.Name == identifier {
					match = t

					return false
				}
			}

			return true
		})

		if match != nil {
			break
		}
	}

	var i, ok = match.Type.(*ast.InterfaceType)

	if !ok {
		panic(match.Type)
	}

	var ms []*MacroMethod

	for _, m := range i.Methods.List {
		var f, ok = m.Type.(*ast.FuncType)

		if !ok {
			panic(m.Type)
		}

		var pfs, rfs []*MacroVar
		var pgs, rgs []*MacroVars

		if f.Params != nil {
			for _, p := range f.Params.List {
				var ns []string

				for _, n := range p.Names {
					ns = append(ns, n.Name)
					pfs = append(pfs, &MacroVar{Name: n.Name, Type: fmt.Sprint(p.Type)})
				}

				pgs = append(pgs, &MacroVars{Names: ns, Type: fmt.Sprint(p.Type)})
			}
		}

		if f.Results != nil {
			for _, r := range f.Results.List {
				var ns []string

				for _, n := range r.Names {
					ns = append(ns, n.Name)
					rfs = append(rfs, &MacroVar{Name: n.Name, Type: fmt.Sprint(r.Type)})
				}

				rgs = append(rgs, &MacroVars{Names: ns, Type: fmt.Sprint(r.Type)})
			}
		}

		ms = append(ms, &MacroMethod{
			Doc:            macroCommentGroup(m.Doc),
			Name:           m.Names[0].Name,
			ParamsFlat:     pfs,
			ParamsGrouped:  pgs,
			ResultsFlat:    rfs,
			ResultsGrouped: rgs,
		})
	}

	var x = &MacroInterface{
		Doc:       macroCommentGroup(match.Doc),
		Methods:   ms,
		Name:      match.Name.Name,
		Package:   package_,
		Qualifier: bp.Name,
	}

	return x, nil
}

type MacroMethod struct {
	Doc            []string
	Name           string
	ParamsFlat     []*MacroVar
	ParamsGrouped  []*MacroVars
	ResultsFlat    []*MacroVar
	ResultsGrouped []*MacroVars
}

type MacroType struct {
	Doc       []string
	Interface *MacroInterface
	Name      string
	Package   string
	Pointer   bool
	Receiver  string
}

type MacroVar struct {
	Name string
	Type string
}

type MacroVars struct {
	Names []string
	Type  string
}
