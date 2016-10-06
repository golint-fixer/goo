package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/willfaught/goo"
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

func macro(p *program) error {
	var bs, err = ioutil.ReadAll(p.in)

	if err != nil {
		return fmt.Errorf("cannot read %v: %v", p.in.Name(), err)
	}

	var done = func() error {
		if _, err := p.out.Write(bs); err != nil {
			return fmt.Errorf("cannot write %v: %v", p.out.Name(), err)
		}

		return nil
	}

	if !p.preprocess {
		return done()
	}

	bs = goo.MacroPreprocess(bs)

	if !p.process {
		return done()
	}

	if bs, err = goo.MacroProcess(p.in.Name(), bs, p.data); err != nil {
		return fmt.Errorf("cannot process macro: %v", err)
	}

	if !p.format {
		return done()
	}

	if bs, err = goo.MacroFormat(bs); err != nil {
		return fmt.Errorf("cannot format macro: %v", err)
	}

	return done()
}

func main() {
	var p = newProgram()

	defer func() {
		p.app.FatalIfError(p.in.Close(), "cannot close %v: ", p.in.Name())
		p.app.FatalIfError(p.out.Close(), "cannot close %v: ", p.out.Name())
	}()

	switch p.cmd {
	case "macro":
		p.app.FatalIfError(macro(p), "")

	default:
		panic(p.cmd)
	}
}

type program struct {
	app        *kingpin.Application
	cmd        string
	data       interface{}
	format     bool
	in         *os.File
	out        *os.File
	preprocess bool
	process    bool
	receiver   string
}

func newProgram() *program {
	var app = kingpin.New(os.Args[0], "The missing link.")
	var macro = app.Command("macro", "Run a macro.")
	var prog = program{app: app}

	app.Author("Will Faught")
	app.Version("0.1.0")
	app.HelpFlag.Short('h')

	var doc = macro.Flag("doc", "Struct doc line.").Short('d').Strings()
	var fields = macro.Flag("field", "Struct field macro data.").Short('f').StringMap()
	var interface_ = macro.Flag("interface", "Struct interface.").Short('n').String()
	var json_ = macro.Flag("json", "JSON macro data. Overrides fields.").Short('j').String()
	var package_ = macro.Flag("package", "Struct package.").Short('p').String()
	var receiver = macro.Flag("receiver", "Struct receiver.").Short('r').String()
	var struct_ = macro.Flag("struct", "Struct name.").Short('s').String()

	macro.Flag("format", "Format the macro.").Default("true").BoolVar(&prog.format)
	macro.Flag("in", "Input file path. Defaults to standard in.").Short('i').FileVar(&prog.in)
	macro.Flag("out", "Output file path. Defaults to standard out.").Short('o').OpenFileVar(&prog.out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	macro.Flag("preprocess", "Preprocess the macro.").Default("true").BoolVar(&prog.preprocess)
	macro.Flag("process", "Process the macro.").Default("true").BoolVar(&prog.process)

	var args = os.Args[1:]
	var cmd = kingpin.MustParse(app.Parse(args))

	if cmd == "" {
		app.Usage(args)
		os.Exit(0)
	}

	prog.cmd = cmd

	if *json_ != "" {
		app.FatalIfError(json.Unmarshal([]byte(*json_), &prog.data), "json is invalid")
	} else if *interface_ != "" {
		var ss = strings.Split(*interface_, ".")
		var pkg = strings.Join(ss[:len(ss)-1], ".")
		var id = ss[len(ss)-1]
		var i, err = goo.MacroInterface(pkg, id)

		if err != nil {
			app.FatalIfError(err, "cannot use interface %v", *interface_)
		}

		prog.data = &goo.Struct{Doc: *doc, Interface: i, Name: *struct_, Package: *package_, Receiver: *receiver}
	} else {
		prog.data = *fields
	}

	convert(prog.data)

	if prog.in == nil {
		prog.in = os.Stdin
	}

	if prog.out == nil {
		prog.out = os.Stdout
	}

	return &prog
}
