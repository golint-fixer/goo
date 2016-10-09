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

var identifier = regexp.MustCompile(`(?m:^\w+$)`)

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
	newProgram().run()
}

type macro struct {
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

func newMacro(app *kingpin.Application) *macro {
	var m = macro{app: app}
	var c = app.Command("macro", "Run a macro.")

	c.Flag("field", "Struct field macro data.").Short('f').StringMapVar(&m.fields)
	c.Flag("format", "Format the macro.").Default("true").BoolVar(&m.format)
	c.Flag("in", "Input file path. Defaults to standard in.").Short('i').FileVar(&m.in)
	c.Flag("json", "JSON macro data. Overrides field.").Short('j').StringVar(&m.json_)
	c.Flag("out", "Output file path. Defaults to standard out.").Short('o').OpenFileVar(&m.out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	c.Flag("preprocess", "Preprocess the macro.").Default("true").BoolVar(&m.preprocess)
	c.Flag("process", "Process the macro.").Default("true").BoolVar(&m.process)

	return &m
}

func (m *macro) run() error {
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

	bs, err = m.pipeline(m.in.Name(), bs, m.data)

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

func (m *macro) pipeline(name string, bs []byte, data interface{}) ([]byte, error) {
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

type program struct {
	app      *kingpin.Application
	macro    *macro
	resource *resource
	stub     *stub
}

func newProgram() *program {
	var app = kingpin.New(os.Args[0], "Fill the gaps.")

	app.Author("Will Faught")
	app.Version("0.1.0")
	app.HelpFlag.Short('h')

	return &program{app: app, macro: newMacro(app), stub: newStub(app), resource: newResource(app)}
}

func (p *program) run() {
	var a = os.Args[1:]

	switch kingpin.MustParse(p.app.Parse(a)) {
	case "macro":
		p.app.FatalIfError(p.macro.run(), "")

	case "stub":
		p.app.FatalIfError(p.stub.run(), "")

	case "resource":
		p.app.FatalIfError(p.resource.run(), "")

	default:
		p.app.Usage(a)
	}
}

type resource struct {
	app      *kingpin.Application
	compress bool
	input    *os.File
	name     string
	output   *os.File
	package_ string
}

func newResource(app *kingpin.Application) *resource {
	var r = resource{app: app}
	var c = app.Command("resource", "Create a resource.")

	c.Arg("input", "Input file.").Required().FileVar(&r.input)

	c.Flag("compress", "Resource compression.").Short('c').BoolVar(&r.compress)
	c.Flag("name", "Resource name.").Short('n').StringVar(&r.name)
	c.Flag("output", "Output file.").Short('o').OpenFileVar(&r.output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	c.Flag("package", "Resource package.").Short('p').StringVar(&r.package_)

	return &r
}

func (r *resource) run() (err error) {
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
		if r.output, err = os.OpenFile(outputNameAbs, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644); err != nil {
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

	var res, ok = goo.Resource("../../internal/resource/resource.go")

	if !ok {
		panic(ok)
	}

	if bs, err = (&macro{format: true, preprocess: true, process: true}).pipeline(inputName, res, data); err != nil {
		return fmt.Errorf("cannot create resource: %v", err)
	}

	if n, err := r.output.Write(bs); err != nil {
		return fmt.Errorf("cannot write %v: %v", outputName, err)
	} else if n != len(bs) {
		return fmt.Errorf("cannot write %v: %v bytes not written", outputName, len(bs)-n)
	}

	return nil
}

type stub struct {
	app        *kingpin.Application
	interface_ string
	name       string
	out        *os.File
	package_   string
	receiver   string
	struct_    *goo.Struct
}

func newStub(app *kingpin.Application) *stub {
	var s = stub{app: app}
	var c = app.Command("stub", "Stub an interface.")

	c.Arg("interface", "Struct interface. Qualified with its package path.").Required().StringVar(&s.interface_)
	c.Flag("name", "Struct name.").Short('s').Default("Stub").StringVar(&s.name)
	c.Flag("out", "Output file path.").Short('o').Default("/dev/stdout").OpenFileVar(&s.out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	c.Flag("package", "Struct package.").Short('p').Default("main").StringVar(&s.package_)
	c.Flag("receiver", "Struct receiver.").Short('r').StringVar(&s.receiver)

	return &s
}

func (s *stub) run() error {
	var ss = strings.Split(s.interface_, ".")
	var pkg = strings.Join(ss[:len(ss)-1], ".")
	var id = ss[len(ss)-1]
	var i, err = goo.MacroInterface(pkg, id)

	if err != nil {
		s.app.FatalIfError(err, "cannot use interface %v", s.interface_)
	}

	s.struct_ = &goo.Struct{Interface: i, Name: s.name, Package: s.package_, Receiver: s.receiver}

	return (&macro{data: s.struct_, in: nil, out: s.out}).run()
}
