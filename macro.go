package goo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/fatih/camelcase"
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
	macroPackage     = regexp.MustCompile(`(?m:^)\w+(?m:$)`)
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
	"PERCENT":     []byte(`%`),
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

// MacroFormat returns input formatted with go/format.
func MacroFormat(input []byte) ([]byte, error) {
	return format.Source(input)
}

// MacroPreprocess returns input with generate comments stripped, identifiers of
// the form __My_ID__ renamed to {{.My_ID}}, and single-line comments beginning
// with a slash and multi-line comments beginning and ending with an asterisk
// replaced with their content.
func MacroPreprocess(input []byte) []byte {
	input = macroGenerate.ReplaceAll(input, macroEmpty)
	input = macroIdentifier.ReplaceAllFunc(input, macroPreprocess)
	input = macroLine.ReplaceAll(input, macroDollarOne)
	input = macroMultiLine.ReplaceAll(input, macroDollarOne)

	return input
}

// MacroProcess returns input ran as a text/template with name and data.
func MacroProcess(name string, input []byte, data interface{}) ([]byte, error) {
	var t, err = template.New(name).Parse(string(input))

	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	if err := t.Execute(&b, data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// MacroRun calls MacroPreprocess, MacroProcess, and then MacroFormat with input
// and returns the result.
func MacroRun(name string, input []byte, data interface{}) ([]byte, error) {
	return (&Macro{Data: data, Name: name}).Run(input)
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

func macroEscapeFunc(bs []byte) []byte {
	return macroPlaceholder
}

func macroEscapeXFunc(bs []byte) []byte {
	return macroPlaceholderX
}

func macroExprString(e ast.Expr) string {
	if e == nil {
		return ""
	}

	switch e := e.(type) {
	case *ast.BasicLit:
		return e.Value

	case *ast.ArrayType:
		return fmt.Sprintf("[%v]%v", macroExprString(e.Len), macroExprString(e.Elt))

	case *ast.ChanType:
		var recv, send string

		if e.Arrow != token.NoPos {
			if e.Dir == ast.RECV {
				recv = "<-"
			} else if e.Dir == ast.SEND {
				send = "<-"
			}
		}

		return fmt.Sprintf("%vchan%v %v", recv, send, macroExprString(e.Value))

	case *ast.Ident:
		return e.Name

	case *ast.Ellipsis:
		return fmt.Sprintf("...%v", macroExprString(e.Elt))

	case *ast.MapType:
		return fmt.Sprintf("map[%v]%v", macroExprString(e.Key), macroExprString(e.Value))

	case *ast.InterfaceType:
		var fields []string

		for _, f := range e.Methods.List {
			var field string

			if len(f.Names) == 0 {
				field = macroExprString(f.Type)
			} else {
				var ns []string

				for _, n := range f.Names {
					ns = append(ns, n.Name)
				}

				field = fmt.Sprintf("%v %v", strings.Join(ns, ", "), macroExprString(f.Type))
			}

			if f.Tag != nil {
				if v := f.Tag.Value; strings.Contains(v, "`") {
					field = fmt.Sprintf(`%v "%v"`, field, v)
				} else {
					field = fmt.Sprintf("%v `%v`", field, v)
				}
			}

			fields = append(fields, field)
		}

		return fmt.Sprintf("interface{%v}", strings.Join(fields, "; "))

	case *ast.StructType:
		var fields []string

		for _, f := range e.Fields.List {
			var field string

			if len(f.Names) == 0 {
				field = macroExprString(f.Type)
			} else {
				var ns []string

				for _, n := range f.Names {
					ns = append(ns, n.Name)
				}

				field = fmt.Sprintf("%v %v", strings.Join(ns, ", "), macroExprString(f.Type))
			}

			if f.Tag != nil {
				if v := f.Tag.Value; strings.Contains(v, "`") {
					field = fmt.Sprintf(`%v "%v"`, field, v)
				} else {
					field = fmt.Sprintf("%v `%v`", field, v)
				}
			}

			fields = append(fields, field)
		}

		return fmt.Sprintf("struct{%v}", strings.Join(fields, "; "))

	default:
		panic(e)
	}
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

type Macro struct {
	Data              interface{}
	DisableFormat     bool
	DisablePreprocess bool
	DisableProcess    bool
	Name              string
}

func (m *Macro) Run(input []byte) ([]byte, error) {
	var err error

	if m.DisablePreprocess {
		return input, nil
	}

	input = MacroPreprocess(input)

	if m.DisableProcess {
		return input, nil
	}

	if input, err = MacroProcess(m.Name, input, m.Data); err != nil {
		return nil, fmt.Errorf("cannot process macro: %v", err)
	}

	if m.DisableFormat {
		return input, nil
	}

	if input, err = MacroFormat(input); err != nil {
		return nil, fmt.Errorf("cannot format macro: %v", err)
	}

	return input, nil
}

type MacroInput struct {
	Macro
	Input string
}

func (m *MacroInput) Run() ([]byte, error) {
	var input, err = ioutil.ReadFile(m.Input)

	if err != nil {
		return nil, fmt.Errorf("cannot read %v: %v", m.Input, err)
	}

	m.Name = m.Input

	return m.Macro.Run(input)
}

type MacroInputOutput struct {
	Macro
	Input  string
	Output string
}

func (m *MacroInputOutput) Run() error {
	var input, err = ioutil.ReadFile(m.Input)

	if err != nil {
		return fmt.Errorf("cannot read %v: %v", m.Input, err)
	}

	m.Name = m.Input

	output, err := m.Macro.Run(input)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(m.Output, output, 0644); err != nil {
		return fmt.Errorf("cannot write %v: %v", m.Output, err)
	}

	return nil
}

type MacroInterface struct {
	InterfaceMethods          []*MacroMethod
	InterfaceName             string
	InterfacePackageName      string
	InterfacePackagePath      string
	InterfacePackageQualifier string
	InterfaceVar              string
	Output                    string
	ReceiverIdentifier        string
	ReceiverTypeName          string
	ReceiverTypePackageName   string
	ReceiverTypePackagePath   string
	ReceiverTypePointer       string
	ReceiverTypeReference     string
}

func (m *MacroInterface) Init() error {
	if m.ReceiverIdentifier == "" {
		var r, err = m.receiver()

		if err != nil {
			return err
		}

		m.ReceiverIdentifier = r
	}

	if m.InterfaceMethods == nil {
		var ms, err = m.methods()

		if err != nil {
			return err
		}

		m.InterfaceMethods = ms
	}

	if m.InterfaceVar == "" {
		var i, _ = utf8.DecodeRuneInString(strings.ToLower(m.InterfaceName))

		m.InterfaceVar = fmt.Sprintf("%c", i)
	}

	if m.Output == "" {
		m.Output = fmt.Sprintf("%v_%v.go", strings.ToLower(m.ReceiverTypeName), strings.ToLower(m.InterfaceName))
	}

	var outputPackageName, outputPackagePath, err = m.filePackage(m.Output)

	if err != nil {
		return err
	}

	if m.ReceiverTypePackageName == "" {
		m.ReceiverTypePackageName = outputPackageName
	}

	if m.ReceiverTypePackagePath == "" {
		m.ReceiverTypePackagePath = outputPackagePath
	}

	if m.InterfacePackageQualifier == "" && m.InterfacePackagePath != outputPackagePath {
		m.InterfacePackageQualifier = m.InterfacePackageName + "."
	}

	return nil
}

func (m *MacroInterface) methods() ([]*MacroMethod, error) {
	var bp, err = build.Import(m.InterfacePackagePath, "", 0)

	if err != nil {
		return nil, fmt.Errorf("cannot find package %v: %v", m.InterfacePackagePath, err)
	}

	if m.InterfacePackageName == "" {
		m.InterfacePackageName = bp.Name
	}

	aps, err := parser.ParseDir(token.NewFileSet(), bp.Dir, nil, parser.ParseComments)

	if err != nil {
		return nil, fmt.Errorf("cannot parse package %v: %v", m.InterfacePackagePath, err)
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

				if t.Name.Name == m.InterfaceName {
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

	var it, ok = match.Type.(*ast.InterfaceType)

	if !ok {
		panic(match.Type)
	}

	var methods []*MacroMethod

	for _, method := range it.Methods.List {
		// TODO: Handle embedded interfaces.
		var f, ok = method.Type.(*ast.FuncType)

		if !ok {
			panic(method.Type)
		}

		var pf, pg = m.vars(f.Params.List)
		var rf, rg = m.vars(f.Results.List)

		methods = append(methods, &MacroMethod{
			Doc:            macroCommentGroup(method.Doc),
			Interface:      m,
			Name:           method.Names[0].Name,
			ParamsFlat:     pf,
			ParamsGrouped:  pg,
			ResultsFlat:    rf,
			ResultsGrouped: rg,
		})
	}

	return methods, nil
}

func (m *MacroInterface) filePackage(file string) (name, path string, err error) {
	abs, err := filepath.Abs(file)

	if err != nil {
		return "", "", fmt.Errorf("cannot get absolute path for %v: %v", file, err)
	}

	var dir = filepath.Dir(abs)

	if p, err := build.ImportDir(dir, 0); err == nil {
		return p.Name, p.ImportPath, nil
	}

	if p := os.Getenv("GOPATH"); p != "" && filepath.IsAbs(p) && filepath.HasPrefix(abs, p) {
		var names = strings.Split(strings.TrimPrefix(abs, p), fmt.Sprintf("%c", filepath.Separator))
		var name = names[len(names)-1]
		var path = strings.Join(names[1:], fmt.Sprintf("%c", filepath.Separator))

		if macroPackage.MatchString(name) {
			return name, path, nil
		}
	}

	if b := filepath.Base(dir); macroPackage.MatchString(b) {
		return b, "", nil
	}

	return "main", "", nil
}

func (m *MacroInterface) receiver() (string, error) {
	var names = map[string]struct{}{}

	for _, method := range m.InterfaceMethods {
		for _, p := range method.ParamsFlat {
			names[p.Name] = struct{}{}
		}

		for _, r := range method.ResultsFlat {
			names[r.Name] = struct{}{}
		}
	}

	var words = camelcase.Split(m.ReceiverTypeName)
	var initials []string

	for i, word := range words {
		var r, _ = utf8.DecodeRuneInString(word)

		if r == utf8.RuneError {
			return "", fmt.Errorf("type %v is invalid: invalid utf8 string", m.ReceiverTypeName)
		}

		initials = append(initials, strings.ToLower(fmt.Sprintf("%c", r)))
		words[i] = strings.ToLower(word)
	}

	var tries = []string{
		initials[0],
		initials[len(initials)-1],
		strings.Join(initials, ""),
		words[0],
		words[len(words)-1],
	}

	for _, name := range tries {
		if _, ok := names[name]; !ok {
			return name, nil
		}
	}

	var name = initials[0]

	for {
		name += "_"

		if _, ok := names[name]; !ok {
			break
		}
	}

	return name, nil
}

func (m *MacroInterface) vars(vars []*ast.Field) ([]*MacroVar, []*MacroVars) {
	var flat []*MacroVar
	var grouped []*MacroVars

	for _, v := range vars {
		var names []string

		if len(v.Names) == 0 {
			flat = append(flat, &MacroVar{Interface: m, Type: macroExprString(v.Type)})
		} else {
			for _, n := range v.Names {
				names = append(names, n.Name)
				flat = append(flat, &MacroVar{Interface: m, Name: n.Name, Type: macroExprString(v.Type)})
			}
		}

		grouped = append(grouped, &MacroVars{Interface: m, Names: names, Type: macroExprString(v.Type)})
	}

	return flat, grouped
}

type MacroMethod struct {
	Doc            []string
	Interface      *MacroInterface
	Name           string
	ParamsFlat     []*MacroVar
	ParamsGrouped  []*MacroVars
	ResultsFlat    []*MacroVar
	ResultsGrouped []*MacroVars
}

type MacroOutput struct {
	Macro
	Output string
}

func (m *MacroOutput) Run(input []byte) error {
	var output, err = m.Macro.Run(input)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(m.Output, output, 0644); err != nil {
		return fmt.Errorf("cannot write %v: %v", m.Output, err)
	}

	return nil
}

type MacroVar struct {
	Interface *MacroInterface
	Name      string
	Type      string
}

type MacroVars struct {
	Interface *MacroInterface
	Names     []string
	Type      string
}
