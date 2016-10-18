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
	"io"
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

func convert(d interface{}) interface{} {
	switch t := d.(type) {
	case map[string]interface{}:
		for k, v := range t {
			t[k] = convert(v)
		}

	case []interface{}:
		for i := range t {
			t[i] = convert(t[i])
		}

	case float64:
		if i := int64(t); float64(i) == t {
			d = i
		}
	}

	return d
}

func main() {
	NewProgram().Run()
}

type Macro struct {
	app        *kingpin.Application
	doc        []string
	data       interface{}
	fields     map[string]string
	format     bool
	in         *os.File
	input      io.Reader
	interface_ string
	json_      string
	out        *os.File
	output     io.Writer
	package_   string
	preprocess bool
	process    bool
	receiver   string
	struct_    string
}

func NewMacro(app *kingpin.Application) *Macro {
	var m = Macro{app: app}
	var c = app.Command("macro", "Run a macro.")

	c.Flag("field", "Type field macro data.").Short('f').StringMapVar(&m.fields)
	c.Flag("format", "Format the macro.").Default("true").BoolVar(&m.format)
	c.Flag("in", "Input file path. Defaults to standard in.").Short('i').FileVar(&m.in)
	c.Flag("json", "JSON macro data. Overrides field.").Short('j').StringVar(&m.json_)
	c.Flag("out", "Output file path. Defaults to standard out.").Short('o').OpenFileVar(&m.out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	c.Flag("preprocess", "Preprocess the macro.").Default("true").BoolVar(&m.preprocess)
	c.Flag("process", "Process the macro.").Default("true").BoolVar(&m.process)

	return &m
}

func (m *Macro) Run() error {
	if m.json_ != "" {
		m.app.FatalIfError(json.Unmarshal([]byte(m.json_), &m.data), "json is invalid")
		convert(m.data)
	} else {
		m.data = m.fields
	}

	if m.in == nil {
		m.in = os.Stdin
	}

	if m.out == nil {
		m.out = os.Stdout
	}

	var bs, err = ioutil.ReadAll(m.in)

	if err != nil {
		return fmt.Errorf("cannot read %v: %v", m.in.Name(), err)
	}

	if bs, err = m.Pipeline(m.in.Name(), bs, m.data); err != nil {
		return fmt.Errorf("cannot run macro: %v", err)
	}

	if n, err := m.out.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", m.out.Name(), err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes unwritten", m.out.Name(), len(bs)-n)
	}

	if err := m.in.Close(); err != nil {
		return fmt.Errorf("cannot close %v: %v", m.in.Name(), err)
	}

	if err := m.out.Close(); err != nil {
		return fmt.Errorf("cannot close %v: %v", m.out.Name(), err)
	}

	return nil
}

func (m *Macro) Pipeline(name string, bs []byte, data interface{}) ([]byte, error) {
	if !m.preprocess {
		return bs, nil
	}

	bs = goo.MacroPreprocess(bs)

	if !m.process {
		return bs, nil
	}

	bs, err := goo.MacroProcess(name, bs, data)

	if err != nil {
		return nil, fmt.Errorf("cannot process macro: %v", err)
	}

	if !m.format {
		return bs, nil
	}

	if bs, err = goo.MacroFormat(bs); err != nil {
		return nil, fmt.Errorf("cannot format macro: %v", err)
	}

	return bs, nil
}

type Mock struct {
	app        *kingpin.Application
	identifier string
	interface_ string
	output     string
	package_   string
	type_      string
	value      bool
}

func NewMock(app *kingpin.Application) *Mock {
	var m = Mock{app: app}
	var c = app.Command("mock", "Mock an interface.")

	c.Arg("type", "Mock receiver type.").Required().StringVar(&m.type_)
	c.Arg("identifier", "Mock receiver identifier.").Required().StringVar(&m.identifier)
	c.Arg("interface", "Mock interface.").Required().StringVar(&m.interface_)

	c.Flag("output", "Mock file.").Short('o').StringVar(&m.output)
	c.Flag("package", "Mock package.").Short('p').StringVar(&m.package_)
	c.Flag("value", "Mock receiver is a value type.").Short('v').BoolVar(&m.value)

	return &m
}

func (s *Mock) Run() error {
	var g = qualified.FindStringSubmatch(s.interface_)

	if len(g) != 3 {
		return fmt.Errorf("interface %v is invalid: %v", s.interface_, g)
	}

	var p, t = g[1], g[2]

	if !identifierExported.MatchString(t) {
		return fmt.Errorf("interface %v is invalid: %v", s.interface_, t)
	}

	var i, err = goo.GetMacroInterface(p, t)

	if err != nil {
		s.app.FatalIfError(err, "cannot parse interface %v", s.interface_)
	}

	if p == "" {
		i.Package = ""
		i.Qualifier = ""
	}

	var outputName string

	if s.output == "" {
		outputName = fmt.Sprintf("%v_%v.go", strings.ToLower(s.type_), strings.ToLower(t))
	} else {
		outputName = s.output
	}

	outputNameAbs, err := filepath.Abs(outputName)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", outputName, err)
	}

	var outputNameAbsDir = filepath.Dir(outputNameAbs)
	var outputPackage string

	if s.package_ == "" {
		if p, err := build.ImportDir(outputNameAbsDir, 0); err == nil {
			s.package_ = p.Name
			outputPackage = p.ImportPath
		} else if b := filepath.Base(outputNameAbsDir); identifier.MatchString(b) {
			s.package_ = b
		} else {
			s.package_ = "main"
		}
	}

	var resource, ok = goo.Resource("../../internal/mock/mock.go")

	if !ok {
		panic(ok)
	}

	if p == outputPackage {
		i.Package = ""
		i.Qualifier = ""
	}

	for _, m := range i.Methods {
		m.Receiver = s.identifier
	}

	var d = &goo.MacroType{Interface: i, Name: s.type_, Package: s.package_, Pointer: !s.value, Receiver: s.identifier}
	bs, err := (&Macro{format: true, preprocess: true, process: true}).Pipeline(outputName, resource, d)

	if err != nil {
		return fmt.Errorf("cannot create mock: %v", err)
	}

	output, err := os.OpenFile(outputNameAbs, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("cannot open %v: %v", outputName, err)
	}

	if n, err := output.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", outputName, err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes not written", outputName, len(bs)-n)
	}

	if err := output.Close(); err != nil {
		return fmt.Errorf("cannot close %v: %v", outputName, err)
	}

	return nil
}

type Program struct {
	App      *kingpin.Application
	Macro    *Macro
	Mock     *Mock
	Resource *Resource
	Stub     *Stub
	Wrap     *Wrap
}

func NewProgram() *Program {
	var app = kingpin.New(os.Args[0], "Fill the gaps.")

	app.Author("Will Faught")
	app.Version("0.1.0")
	app.HelpFlag.Short('h')

	return &Program{
		App:      app,
		Macro:    NewMacro(app),
		Mock:     NewMock(app),
		Stub:     NewStub(app),
		Resource: NewResource(app),
		Wrap:     NewWrap(app),
	}
}

func (p *Program) Run() {
	var a = os.Args[1:]

	switch kingpin.MustParse(p.App.Parse(a)) {
	case "macro":
		p.App.FatalIfError(p.Macro.Run(), "")

	case "mock":
		p.App.FatalIfError(p.Mock.Run(), "")

	case "stub":
		p.App.FatalIfError(p.Stub.Run(), "")

	case "resource":
		p.App.FatalIfError(p.Resource.Run(), "")

	case "wrap":
		p.App.FatalIfError(p.Wrap.Run(), "")

	default:
		p.App.Usage(a)
	}
}

type Resource struct {
	app      *kingpin.Application
	compress bool
	input    *os.File
	name     string
	output   *os.File
	package_ string
}

func NewResource(app *kingpin.Application) *Resource {
	var r = Resource{app: app}
	var c = app.Command("resource", "Create a resource.")

	c.Arg("input", "Input file.").Required().FileVar(&r.input)

	c.Flag("compress", "Resource compression.").Short('c').BoolVar(&r.compress)
	c.Flag("name", "Resource name.").Short('n').StringVar(&r.name)
	c.Flag("output", "Output file.").Short('o').OpenFileVar(&r.output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	c.Flag("package", "Resource package.").Short('p').StringVar(&r.package_)

	return &r
}

func (r *Resource) Run() (err error) {
	defer func() {
		if err2 := r.input.Close(); err2 != nil && err == nil {
			err = fmt.Errorf("cannot close %v: %v", r.input.Name(), err2)
		}

		if r.output != nil {
			if err2 := r.output.Close(); err2 != nil && err == nil {
				err = fmt.Errorf("cannot close %v: %v", r.output.Name(), err2)
			}
		}
	}()

	var inputName = r.input.Name()
	inputNameAbs, err := filepath.Abs(inputName)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", inputName, err)
	}

	var outputName string

	if r.output == nil {
		outputName = filepath.Base(inputNameAbs)

		if e := filepath.Ext(outputName); e == "" {
			outputName += ".go"
		} else if e != ".go" {
			outputName = strings.TrimSuffix(outputName, e) + ".go"
		}
	} else {
		outputName = r.output.Name()
	}

	outputNameAbs, err := filepath.Abs(outputName)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", outputName, err)
	}

	var outputNameAbsDir = filepath.Dir(outputNameAbs)

	if r.name == "" {
		if r.name, err = filepath.Rel(outputNameAbsDir, inputNameAbs); err != nil {
			return fmt.Errorf("cannot get relative path from %v to %v: %v", outputNameAbsDir, inputNameAbs, err)
		}
	}

	if r.package_ == "" {
		if p, err := build.ImportDir(outputNameAbsDir, 0); err == nil {
			r.package_ = p.Name
		} else if b := filepath.Base(outputNameAbsDir); identifier.MatchString(b) {
			r.package_ = b
		} else {
			r.package_ = "main"
		}
	}

	if r.output == nil {
		if r.output, err = os.OpenFile(outputNameAbs, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
			return fmt.Errorf("cannot open %v: %v", outputName, err)
		}
	}

	bs, err := ioutil.ReadAll(r.input)

	if err != nil {
		return fmt.Errorf("cannot read %v: %v", inputName, err)
	}

	if r.compress {
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
		Compressed: fmt.Sprint(r.compress),
		Data:       fmt.Sprintf("%#v", bs),
		Name:       r.name,
		Package:    r.package_,
	}

	var resource, ok = goo.Resource("../../internal/resource/resource.go")

	if !ok {
		panic(ok)
	}

	if bs, err = (&Macro{format: true, preprocess: true, process: true}).Pipeline(inputName, resource, data); err != nil {
		return fmt.Errorf("cannot create resource: %v", err)
	}

	if n, err := r.output.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", outputName, err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes not written", outputName, len(bs)-n)
	}

	return nil
}

type Stub struct {
	app        *kingpin.Application
	identifier string
	interface_ string
	output     string
	package_   string
	type_      string
	value      bool
}

func NewStub(app *kingpin.Application) *Stub {
	var s = Stub{app: app}
	var c = app.Command("stub", "Stub an interface.")

	c.Arg("type", "Stub receiver type.").Required().StringVar(&s.type_)
	c.Arg("interface", "Stub interface.").Required().StringVar(&s.interface_)

	c.Flag("identifier", "Stub receiver identifier.").Short('i').StringVar(&s.identifier)
	c.Flag("output", "Stub file.").Short('o').StringVar(&s.output)
	c.Flag("package", "Stub package.").Short('p').StringVar(&s.package_)
	c.Flag("value", "Stub receiver is a value type.").Short('v').BoolVar(&s.value)

	return &s
}

func (s *Stub) Run() error {
	var g = qualified.FindStringSubmatch(s.interface_)

	if len(g) != 3 {
		return fmt.Errorf("interface %v is invalid: %v", s.interface_, g)
	}

	var p, t = g[1], g[2]

	if !identifierExported.MatchString(t) {
		return fmt.Errorf("interface %v is invalid: %v", s.interface_, t)
	}

	var i, err = goo.GetMacroInterface(p, t)

	if err != nil {
		s.app.FatalIfError(err, "cannot parse interface %v", s.interface_)
	}

	if p == "" {
		i.Package = ""
		i.Qualifier = ""
	}

	var outputName string

	if s.output == "" {
		outputName = fmt.Sprintf("%v_%v.go", strings.ToLower(s.type_), strings.ToLower(t))
	} else {
		outputName = s.output
	}

	outputNameAbs, err := filepath.Abs(outputName)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", outputName, err)
	}

	var outputNameAbsDir = filepath.Dir(outputNameAbs)
	var outputPackage string

	if s.package_ == "" {
		if p, err := build.ImportDir(outputNameAbsDir, 0); err == nil {
			s.package_ = p.Name
			outputPackage = p.ImportPath
		} else if b := filepath.Base(outputNameAbsDir); identifier.MatchString(b) {
			s.package_ = b
		} else {
			s.package_ = "main"
		}
	}

	var resource, ok = goo.Resource("../../internal/stub/stub.go")

	if !ok {
		panic(ok)
	}

	if p == outputPackage {
		i.Package = ""
		i.Qualifier = ""
	}

	var d = &goo.MacroType{Interface: i, Name: s.type_, Package: s.package_, Pointer: !s.value, Receiver: s.identifier}
	bs, err := (&Macro{format: true, preprocess: true, process: true}).Pipeline(outputName, resource, d)

	if err != nil {
		return fmt.Errorf("cannot create stub: %v", err)
	}

	output, err := os.OpenFile(outputNameAbs, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("cannot open %v: %v", outputName, err)
	}

	if n, err := output.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", outputName, err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes not written", outputName, len(bs)-n)
	}

	if err := output.Close(); err != nil {
		return fmt.Errorf("cannot close %v: %v", outputName, err)
	}

	return nil
}

type Wrap struct {
	app        *kingpin.Application
	identifier string
	interface_ string
	output     string
	package_   string
	type_      string
	value      bool
}

func NewWrap(app *kingpin.Application) *Wrap {
	var m = Wrap{app: app}
	var c = app.Command("wrap", "Wrap an interface.")

	c.Arg("type", "Wrap receiver type.").Required().StringVar(&m.type_)
	c.Arg("identifier", "Wrap receiver identifier.").Required().StringVar(&m.identifier)
	c.Arg("interface", "Wrap interface.").Required().StringVar(&m.interface_)

	c.Flag("output", "Wrap file.").Short('o').StringVar(&m.output)
	c.Flag("package", "Wrap package.").Short('p').StringVar(&m.package_)
	c.Flag("value", "Wrap receiver is a value type.").Short('v').BoolVar(&m.value)

	return &m
}

func (s *Wrap) Run() error {
	var g = qualified.FindStringSubmatch(s.interface_)

	if len(g) != 3 {
		return fmt.Errorf("interface %v is invalid: %v", s.interface_, g)
	}

	var p, t = g[1], g[2]

	if !identifierExported.MatchString(t) {
		return fmt.Errorf("interface %v is invalid: %v", s.interface_, t)
	}

	var i, err = goo.GetMacroInterface(p, t)

	if err != nil {
		s.app.FatalIfError(err, "cannot parse interface %v", s.interface_)
	}

	if p == "" {
		i.Package = ""
		i.Qualifier = ""
	}

	var outputName string

	if s.output == "" {
		outputName = fmt.Sprintf("%v_%v.go", strings.ToLower(s.type_), strings.ToLower(t))
	} else {
		outputName = s.output
	}

	outputNameAbs, err := filepath.Abs(outputName)

	if err != nil {
		return fmt.Errorf("cannot get absolute path for %v: %v", outputName, err)
	}

	var outputNameAbsDir = filepath.Dir(outputNameAbs)
	var outputPackage string

	if s.package_ == "" {
		if p, err := build.ImportDir(outputNameAbsDir, 0); err == nil {
			s.package_ = p.Name
			outputPackage = p.ImportPath
		} else if b := filepath.Base(outputNameAbsDir); identifier.MatchString(b) {
			s.package_ = b
		} else {
			s.package_ = "main"
		}
	}

	var resource, ok = goo.Resource("../../internal/wrap/wrap.go")

	if !ok {
		panic(ok)
	}

	if p == outputPackage {
		i.Package = ""
		i.Qualifier = ""
	}

	for _, m := range i.Methods {
		m.Receiver = s.identifier
	}

	var d = &goo.MacroType{Interface: i, Name: s.type_, Package: s.package_, Pointer: !s.value, Receiver: s.identifier}
	bs, err := (&Macro{format: true, preprocess: true, process: true}).Pipeline(outputName, resource, d)

	if err != nil {
		return fmt.Errorf("cannot create wrapper: %v", err)
	}

	output, err := os.OpenFile(outputNameAbs, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("cannot open %v: %v", outputName, err)
	}

	if n, err := output.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", outputName, err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes not written", outputName, len(bs)-n)
	}

	if err := output.Close(); err != nil {
		return fmt.Errorf("cannot close %v: %v", outputName, err)
	}

	return nil
}
