package goo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode/utf8"
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

var (
	macroTransformerIdentifier         = regexp.MustCompile(`(?m:^)\w+(?m:$)`)
	macroTransformerIdentifierExported = regexp.MustCompile(`(?m:^)[A-Z]\w*(?m:$)`)
	macroTransformerQualified          = regexp.MustCompile(`(?m:^)(?:([\w.]+(?:\/[\w.]+)*)\.)?([\w.]+)(?m:$)`)
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

func macroCommentGroupStrings(cg *ast.CommentGroup) []string {
	if cg == nil {
		return nil
	}

	var cs []string

	for _, d := range cg.List {
		cs = append(cs, d.Text)
	}

	return cs
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

type MacroInterface struct {
	Doc       []string
	Initial   string
	Methods   []*MacroMethod
	Name      string
	Package   string
	Qualifier string
}

func NewMacroInterface(interfacePackage, interfaceName string) (*MacroInterface, error) {
	var bp, err = build.Import(interfacePackage, "", 0)

	if err != nil {
		return nil, fmt.Errorf("cannot find package %v: %v", interfacePackage, err)
	}

	aps, err := parser.ParseDir(token.NewFileSet(), bp.Dir, nil, parser.ParseComments)

	if err != nil {
		return nil, fmt.Errorf("cannot parse package %v: %v", interfacePackage, err)
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

				if t.Name.Name == interfaceName {
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

	var initial, _ = utf8.DecodeRuneInString(strings.ToLower(match.Name.Name))

	var mi = &MacroInterface{
		Doc:       macroCommentGroupStrings(match.Doc),
		Initial:   fmt.Sprintf("%c", initial),
		Name:      match.Name.Name,
		Package:   interfacePackage,
		Qualifier: bp.Name,
	}

	for _, m := range it.Methods.List {
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
					pfs = append(pfs, &MacroVar{Name: n.Name, Type: macroExprString(p.Type)})
				}

				pgs = append(pgs, &MacroVars{Names: ns, Type: macroExprString(p.Type)})
			}
		}

		if f.Results != nil {
			for _, r := range f.Results.List {
				var ns []string

				if len(r.Names) == 0 {
					rfs = append(rfs, &MacroVar{Type: macroExprString(r.Type)})
				} else {
					for _, n := range r.Names {
						ns = append(ns, n.Name)
						rfs = append(rfs, &MacroVar{Name: n.Name, Type: macroExprString(r.Type)})
					}
				}

				rgs = append(rgs, &MacroVars{Names: ns, Type: macroExprString(r.Type)})
			}
		}

		mi.Methods = append(mi.Methods, &MacroMethod{
			Doc:            macroCommentGroupStrings(m.Doc),
			Interface:      mi,
			Name:           m.Names[0].Name,
			ParamsFlat:     pfs,
			ParamsGrouped:  pgs,
			ResultsFlat:    rfs,
			ResultsGrouped: rgs,
		})
	}

	return mi, nil
}

type MacroMethod struct {
	Doc            []string
	Interface      *MacroInterface
	Name           string
	ParamsFlat     []*MacroVar
	ParamsGrouped  []*MacroVars
	Receiver       string
	ResultsFlat    []*MacroVar
	ResultsGrouped []*MacroVars
}

type MacroTransformer struct {
	Identifier  string
	Interface   string
	Output      string
	Package     string
	SkipFormat  bool
	SkipProcess bool
	Type        string
	Value       bool
}

func (t *MacroTransformer) Transform(input []byte) error {
	var s = macroTransformerQualified.FindStringSubmatch(t.Interface)

	if len(s) != 3 {
		return fmt.Errorf("type %v is invalid", t.Interface)
	}

	var qualifiedPackage, qualifiedType = s[1], s[2]

	if !macroTransformerIdentifierExported.MatchString(qualifiedType) {
		return fmt.Errorf("type %v is invalid", qualifiedType)
	}

	var mi, err = NewMacroInterface(qualifiedPackage, qualifiedType)

	if err != nil {
		return fmt.Errorf("cannot parse interface %v: %v", t.Interface, err)
	}

	if t.Output == "" {
		t.Output = fmt.Sprintf("%v_%v.go", strings.ToLower(t.Type), strings.ToLower(qualifiedType))
	}

	outputAbs, err := filepath.Abs(t.Output)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", t.Output, err)
	}

	var outputAbsDir = filepath.Dir(outputAbs)
	var packageName, packagePath string

	if p, err := build.ImportDir(outputAbsDir, 0); err == nil {
		packageName = p.Name
		packagePath = p.ImportPath
	} else if b := filepath.Base(outputAbsDir); macroTransformerIdentifier.MatchString(b) {
		packageName = b
	} else {
		packageName = "main"
	}

	if t.Package == "" {
		t.Package = packageName
	}

	if qualifiedPackage == "" || qualifiedPackage == packagePath {
		mi.Package = ""
		mi.Qualifier = ""
	}

	for _, m := range mi.Methods {
		m.Receiver = t.Identifier
	}

	input = MacroPreprocess(input)

	if t.SkipProcess {
		return t.writeFile(t.Output, input)
	}

	input, err = MacroProcess(t.Output, input, &MacroType{
		Interface: mi,
		Name:      t.Type,
		Package:   t.Package,
		Pointer:   !t.Value,
		Receiver:  t.Identifier,
	})

	if err != nil {
		return fmt.Errorf("cannot process macro: %v", err)
	}

	if t.SkipFormat {
		return t.writeFile(t.Output, input)
	}

	input, err = MacroFormat(input)

	if err != nil {
		return fmt.Errorf("cannot format macro: %v", err)
	}

	return t.writeFile(t.Output, input)
}

func (t *MacroTransformer) writeFile(path string, bs []byte) error {
	var file, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("cannot open %v: %v", path, err)
	}

	if n, err := file.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", path, err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes not written", path, len(bs)-n)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("cannot close %v: %v", path, err)
	}

	return nil
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
