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

var identifier = regexp.MustCompile(`(?m:^)\w+(?m:$)`)

func main() {
	var app = kingpin.New(os.Args[0], "Fill the gaps.")

	app.Author("Will Faught")
	app.Version("0.1.0")
	app.HelpFlag.Short('h')

	var args = os.Args[1:]
	var err error

	var (
		macro    = NewGeneralMacro(app.Command("macro", "Run a macro."))
		mock     = NewInterfaceMacro(app.Command("mock", "Mock an interface."))
		resource = NewResourceMacro(app.Command("resource", "Add a resource."))
		stub     = NewInterfaceMacro(app.Command("stub", "Stub an interface."))
		wrap     = NewInterfaceMacro(app.Command("wrap", "Wrap an interface."))
	)

	switch kingpin.MustParse(app.Parse(args)) {
	case "macro":
		err = macro.Run()

	case "mock":
		err = mock.Run(goo.Resource("github.com/willfaught/goo/cmd/goo", "../../internal/mock/mock.go"))

	case "stub":
		err = stub.Run(goo.Resource("github.com/willfaught/goo/cmd/goo", "../../internal/stub/stub.go"))

	case "resource":
		err = resource.Run()

	case "wrap":
		err = wrap.Run(goo.Resource("github.com/willfaught/goo/cmd/goo", "../../internal/wrap/wrap.go"))

	default:
		app.Usage(args)
	}

	app.FatalIfError(err, "")
}

type GeneralMacro struct {
	goo.MacroInputOutput
	Data map[string]string
	JSON string
}

func NewGeneralMacro(c *kingpin.CmdClause) *GeneralMacro {
	var m GeneralMacro

	c.Arg("input", "Input file path.").Required().StringVar(&m.Input)
	c.Arg("output", "Output file path.").Required().StringVar(&m.Output)

	c.Flag("data", "Macro data map as JSON.").Short('d').StringVar(&m.JSON)
	c.Flag("disable-format", "Do not format the macro output.").BoolVar(&m.DisableFormat)
	c.Flag("disable-preprocess", "Do not preprocess the macro output.").BoolVar(&m.DisablePreprocess)
	c.Flag("disable-process", "Do not process the macro output.").BoolVar(&m.DisableProcess)
	c.Flag("key", "Macro data key and value.").Short('k').StringMapVar(&m.Data)

	return &m
}

func (m *GeneralMacro) Run() error {
	if m.JSON == "" {
		m.MacroInputOutput.Data = m.Data
	} else if err := json.Unmarshal([]byte(m.JSON), &m.MacroInputOutput.Data); err != nil {
		return fmt.Errorf("json invalid: %v", err)
	}

	return m.MacroInputOutput.Run()
}

type InterfaceMacro struct {
	goo.MacroInterface
	goo.MacroOutput
	Value bool
}

func NewInterfaceMacro(c *kingpin.CmdClause) *InterfaceMacro {
	var m InterfaceMacro

	c.Arg("interface-package-path", "Interface package path.").Required().StringVar(&m.InterfacePackagePath)
	c.Arg("interface-name", "Interface name.").Required().StringVar(&m.InterfaceName)
	c.Arg("receiver-type-name", "Receiver type name.").Required().StringVar(&m.ReceiverTypeName)

	c.Flag("disable-format", "Do not format the macro output.").BoolVar(&m.DisableFormat)
	c.Flag("disable-preprocess", "Do not preprocess the macro output.").BoolVar(&m.DisablePreprocess)
	c.Flag("disable-process", "Do not process the macro output.").BoolVar(&m.DisableProcess)
	c.Flag("interface-package-name", "Interface package name.").StringVar(&m.InterfacePackageName)
	c.Flag("output", "Output file path.").Short('o').StringVar(&m.MacroInterface.Output)
	c.Flag("receiver-identifier", "Receiver identifier.").StringVar(&m.ReceiverIdentifier)
	c.Flag("receiver-type-package-name", "Receiver type package name.").StringVar(&m.ReceiverTypePackageName)
	c.Flag("receiver-type-value", "Receiver type is a value.").BoolVar(&m.Value)

	return &m
}

func (m *InterfaceMacro) Run(input []byte) error {
	if err := m.MacroInterface.Init(); err != nil {
		return err
	}

	m.Data = m

	if m.MacroOutput.Output == "" {
		m.MacroOutput.Output = m.MacroInterface.Output
	}

	m.Name = m.MacroOutput.Output

	if !m.Value {
		m.ReceiverTypePointer = "*"
		m.ReceiverTypeReference = "&"
	}

	return m.MacroOutput.Run(input)
}

type ResourceMacro struct {
	goo.MacroOutput
	Compress    bool
	Data        string
	Input       string
	Name        string
	PackageName string
	PackagePath string
	Value       bool
}

func NewResourceMacro(c *kingpin.CmdClause) *ResourceMacro {
	var m ResourceMacro

	c.Arg("input", "Input file path.").Required().StringVar(&m.Input)

	c.Flag("compress", "Resource compression.").Short('c').BoolVar(&m.Compress)
	c.Flag("disable-format", "Do not format the macro output.").BoolVar(&m.DisableFormat)
	c.Flag("disable-preprocess", "Do not preprocess the macro output.").BoolVar(&m.DisablePreprocess)
	c.Flag("disable-process", "Do not process the macro output.").BoolVar(&m.DisableProcess)
	c.Flag("name", "Resource name.").Short('n').StringVar(&m.Name)
	c.Flag("output", "Output file path.").Short('o').StringVar(&m.Output)
	c.Flag("package-name", "Resource package name.").StringVar(&m.PackageName)
	c.Flag("package-path", "Resource package path.").StringVar(&m.PackagePath)

	return &m
}

func (m *ResourceMacro) Run() error {
	var input, err = ioutil.ReadFile(m.Input)

	if err != nil {
		return fmt.Errorf("cannot read %v: %v", m.Input, err)
	}

	inputAbs, err := filepath.Abs(m.Input)

	if err != nil {
		return err
	}

	if m.Output == "" {
		m.Output = filepath.Base(inputAbs)

		if e := filepath.Ext(m.Output); e == "" {
			m.Output += ".go"
		} else if e != ".go" {
			m.Output = strings.TrimSuffix(m.Output, e) + ".go"
		}
	}

	outputAbs, err := filepath.Abs(m.Output)

	if err != nil {
		return err
	}

	var outputAbsDir = filepath.Dir(outputAbs)

	if m.Name == "" {
		if m.Name, err = filepath.Rel(outputAbsDir, inputAbs); err != nil {
			return err
		}
	}

	var packageName, packagePath = m.dirPackage(outputAbsDir)

	if m.PackageName == "" {
		m.PackageName = packageName
	}

	if m.PackagePath == "" {
		m.PackagePath = packagePath
	}

	if m.Compress {
		var b bytes.Buffer
		var w = gzip.NewWriter(&b)

		if n, err := w.Write(input); err != nil {
			panic(err)
		} else if n != len(input) {
			panic(n)
		}

		if err := w.Close(); err != nil {
			panic(err)
		}

		input = b.Bytes()
	}

	m.Data = fmt.Sprintf("%#v", input)
	m.MacroOutput.Data = m

	return m.MacroOutput.Run(goo.Resource("github.com/willfaught/goo/cmd/goo", "../../internal/resource/resource.go"))
}

func (m *ResourceMacro) dirPackage(path string) (packageName string, packagePath string) {
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
