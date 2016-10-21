//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo resource ../../internal/mock/mock.go
//go:generate goo resource ../../internal/resource/resource.go
//go:generate goo resource ../../internal/stub/stub.go
//go:generate goo resource ../../internal/wrap/wrap.go

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/willfaught/goo"
)

var (
	identifier         = regexp.MustCompile(`(?m:^)\w+(?m:$)`)
	identifierExported = regexp.MustCompile(`(?m:^)[A-Z]\w*(?m:$)`)
	qualified          = regexp.MustCompile(`(?m:^)(?:([\w.]+(?:\/[\w.]+)*)\.)?([\w.]+)(?m:$)`)
)

func Convert(d interface{}) interface{} {
	switch t := d.(type) {
	case map[string]interface{}:
		for k, v := range t {
			t[k] = Convert(v)
		}

	case []interface{}:
		for i := range t {
			t[i] = Convert(t[i])
		}

	case float64:
		if i := int64(t); float64(i) == t {
			d = i
		}
	}

	return d
}

func DirPackage(path string) (packageName string, packagePath string) {
	if p, err := build.ImportDir(path, 0); err == nil {
		packageName = p.Name
		packagePath = p.ImportPath
	} else if b := filepath.Base(path); identifier.MatchString(b) {
		packageName = b
	} else {
		packageName = "main"
	}

	return
}

func ParseQualifiedType(q string) (string, string, error) {
	var g = qualified.FindStringSubmatch(q)

	if len(g) != 3 {
		return "", "", fmt.Errorf("type %v is invalid", q)
	}

	var p, t = g[1], g[2]

	if !identifierExported.MatchString(t) {
		return "", "", fmt.Errorf("type %v is invalid", t)
	}

	return p, t, nil
}

func Pipeline(preprocess, process bool, name string, bs []byte, data interface{}) ([]byte, error) {
	bs = goo.MacroPreprocess(bs)

	if preprocess {
		return bs, nil
	}

	bs, err := goo.MacroProcess(name, bs, data)

	if err != nil {
		return nil, fmt.Errorf("cannot process macro: %v", err)
	}

	if process {
		return bs, nil
	}

	if bs, err = goo.MacroFormat(bs); err != nil {
		return nil, fmt.Errorf("cannot format macro: %v", err)
	}

	return bs, nil
}

func ReadFile(path string) ([]byte, error) {
	var file, err = os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("cannot open %v: %v", path, err)
	}

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("cannot read %v: %v", path, err)
	}

	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("cannot close %v: %v", path, err)
	}

	return bs, nil
}

func WriteFile(path string, bs []byte) error {
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

func main() {
	var app = kingpin.New(os.Args[0], "Fill the gaps.")

	app.Author("Will Faught")
	app.Version("0.1.0")
	app.HelpFlag.Short('h')

	var args = os.Args[1:]
	var err error

	var (
		macro    = NewMacro(app)
		mock     = NewTransformer(app, "mock", "Mock", "Mock", "../../internal/mock/mock.go")
		stub     = NewTransformer(app, "stub", "Stub", "Stub", "../../internal/stub/stub.go")
		resource = NewResource(app)
		wrap     = NewTransformer(app, "wrap", "Wrap", "Wrap", "../../internal/wrap/wrap.go")
	)

	switch kingpin.MustParse(app.Parse(args)) {
	case "macro":
		err = macro.Run()

	case "mock":
		err = mock.Run()

	case "stub":
		err = stub.Run()

	case "resource":
		err = resource.Run()

	case "wrap":
		err = wrap.Run()

	default:
		app.Usage(args)
	}

	app.FatalIfError(err, "")
}

type Macro struct {
	Fields     map[string]string
	Input      string
	JSON       string
	Output     string
	Preprocess bool
	Process    bool
}

func NewMacro(app *kingpin.Application) *Macro {
	var m Macro
	var c = app.Command("macro", "Run a macro.")

	c.Arg("input", "Input file.").Required().StringVar(&m.Input)
	c.Arg("output", "Output file.").Required().StringVar(&m.Output)

	c.Flag("field", "Macro data JSON object field.").Short('f').StringMapVar(&m.Fields)
	c.Flag("json", "Macro data JSON.").Short('j').StringVar(&m.JSON)
	c.Flag("preprocess", "Preprocess the macro then stop.").BoolVar(&m.Preprocess)
	c.Flag("process", "Process the macro then stop.").BoolVar(&m.Process)

	return &m
}

func (m *Macro) Run() error {
	var data interface{}

	if m.JSON != "" {
		if err := json.Unmarshal([]byte(m.JSON), &data); err != nil {
			return fmt.Errorf("json is invalid: %v", err)
		}

		Convert(&data)
	} else {
		data = m.Fields
	}

	var bs, err = ReadFile(m.Input)

	if err != nil {
		return err
	}

	if bs, err = Pipeline(m.Preprocess, m.Process, m.Input, bs, data); err != nil {
		return fmt.Errorf("cannot run macro: %v", err)
	}

	return WriteFile(m.Output, bs)
}

type Resource struct {
	Compress bool
	File     string
	Name     string
	Output   string
	Package  string
}

func NewResource(app *kingpin.Application) *Resource {
	var r Resource
	var c = app.Command("resource", "Create a resource.")

	c.Arg("file", "Input file.").Required().StringVar(&r.File)

	c.Flag("compress", "Resource compression.").Short('c').BoolVar(&r.Compress)
	c.Flag("name", "Resource name.").Short('n').StringVar(&r.Name)
	c.Flag("output", "Output file.").Short('o').StringVar(&r.Output)
	c.Flag("package", "Resource package.").Short('p').StringVar(&r.Package)

	return &r
}

func (r *Resource) Run() error {
	var fileAbs, err = filepath.Abs(r.File)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", r.File, err)
	}

	if r.Output == "" {
		r.Output = filepath.Base(fileAbs)

		if e := filepath.Ext(r.Output); e == "" {
			r.Output += ".go"
		} else if e != ".go" {
			r.Output = strings.TrimSuffix(r.Output, e) + ".go"
		}
	}

	outputAbs, err := filepath.Abs(r.Output)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", r.Output, err)
	}

	var outputAbsDir = filepath.Dir(outputAbs)

	if r.Name == "" {
		if r.Name, err = filepath.Rel(outputAbsDir, fileAbs); err != nil {
			return fmt.Errorf("cannot get relative path from %v to %v: %v", outputAbsDir, fileAbs, err)
		}
	}

	if r.Package == "" {
		r.Package, _ = DirPackage(outputAbsDir)
	}

	bs, err := ReadFile(r.File)

	if err != nil {
		return err
	}

	if r.Compress {
		var compressed bytes.Buffer
		var w = gzip.NewWriter(&compressed)

		if n, err := w.Write(bs); err != nil {
			panic(err)
		} else if n != len(bs) {
			panic(n)
		}

		if err := w.Close(); err != nil {
			panic(err)
		}

		bs = compressed.Bytes()
	}

	var data = struct {
		Compressed, Data, Name, Package string
	}{
		Compressed: fmt.Sprint(r.Compress),
		Data:       fmt.Sprintf("%#v", bs),
		Name:       r.Name,
		Package:    r.Package,
	}

	var resource, ok = goo.Resource("../../internal/resource/resource.go")

	if !ok {
		panic(ok)
	}

	if bs, err = Pipeline(false, false, r.File, resource, data); err != nil {
		return fmt.Errorf("cannot create resource: %v", err)
	}

	return WriteFile(r.Output, bs)
}

type Transformer struct {
	Identifier string
	Interface  string
	Output     string
	Package    string
	Resource   string
	Type       string
	Value      bool
}

func NewTransformer(a *kingpin.Application, command, verb, noun, resource string) *Transformer {
	var t = Transformer{Resource: resource}
	var c = a.Command(command, fmt.Sprintf("%v an interface.", verb))

	c.Arg("receiver-type", fmt.Sprintf("%v receiver type.", noun)).Required().StringVar(&t.Type)
	c.Arg("receiver-identifier", fmt.Sprintf("%v receiver identifiet.", noun)).Required().StringVar(&t.Identifier)
	c.Arg("interface", fmt.Sprintf("%v interface.", noun)).Required().StringVar(&t.Interface)

	c.Flag("output", fmt.Sprintf("%v file.", noun)).Short('o').StringVar(&t.Output)
	c.Flag("package", fmt.Sprintf("%v package.", noun)).Short('p').StringVar(&t.Package)
	c.Flag("value", fmt.Sprintf("%v receiver is a value type.", noun)).Short('v').BoolVar(&t.Value)

	return &t
}

func (t *Transformer) Run() error {
	var m = qualified.FindStringSubmatch(t.Interface)

	if len(m) != 3 {
		return fmt.Errorf("type %v is invalid", t.Interface)
	}

	var qp, qt = m[1], m[2]

	if !identifierExported.MatchString(qt) {
		return fmt.Errorf("type %v is invalid", qt)
	}

	var mi, err = goo.GetMacroInterface(qp, qt)

	if err != nil {
		return fmt.Errorf("cannot parse interface %v: %v", t.Interface, err)
	}

	if t.Output == "" {
		t.Output = fmt.Sprintf("%v_%v.go", strings.ToLower(t.Type), strings.ToLower(qt))
	}

	outputPathAbs, err := filepath.Abs(t.Output)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", t.Output, err)
	}

	var dn, dp = DirPackage(filepath.Dir(outputPathAbs))

	if t.Package == "" {
		t.Package = dn
	}

	if qp == "" || qp == dp {
		mi.Package = ""
		mi.Qualifier = ""
	}

	for _, m := range mi.Methods {
		m.Receiver = t.Identifier
	}

	var resource, ok = goo.Resource(t.Resource)

	if !ok {
		panic(ok)
	}

	bs, err := Pipeline(false, false, t.Output, resource, &goo.MacroType{
		Interface: mi,
		Name:      t.Type,
		Package:   t.Package,
		Pointer:   !t.Value,
		Receiver:  t.Identifier,
	})

	if err != nil {
		return fmt.Errorf("cannot create mock: %v", err)
	}

	return WriteFile(t.Output, bs)
}
