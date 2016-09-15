package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

type vars map[string][]string

func (v vars) Set(s string) error {
	var ss = strings.Split(s, "=")

	if len(ss) != 2 {
		return fmt.Errorf("variable %v is invalid", s)
	}

	v[ss[0]] = append(v[ss[0]], ss[1])

	return nil
}

func (v vars) String() string {
	return fmt.Sprint(map[string][]string(v))
}

func main() {
	var v = vars{}
	var pathIn = flag.String("in", "", "Input file path")
	var pathOut = flag.String("out", "", "Output file path")

	flag.Var(&v, "var", "Variable of the form name=value")
	flag.Parse()

	if *pathIn == "" || *pathOut == "" || len(v) == 0 {
		flag.PrintDefaults()
		os.Exit(1)

		return
	}

	var stringType = reflect.TypeOf("")
	var stringSliceType = reflect.SliceOf(stringType)
	var fs []reflect.StructField
	var vs []reflect.Value

	for name, values := range v {
		if len(values) == 1 {
			fs = append(fs, reflect.StructField{Name: name, Type: stringType})
			vs = append(vs, reflect.ValueOf(values[0]))
		} else if len(values) > 1 {
			fs = append(fs, reflect.StructField{Name: name, Type: stringSliceType})
			vs = append(vs, reflect.ValueOf(values))
		}
	}

	var structValue = reflect.New(reflect.StructOf(fs))

	for i := range fs {
		structValue.Elem().Field(i).Set(vs[i])
	}

	var bs, err = ioutil.ReadFile(*pathIn)

	if err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	bs = regexp.MustCompile(`__([\w_]+)__`).ReplaceAll(bs, []byte("{{.$1}}"))
	bs = regexp.MustCompile(`\/\/\/\s*(\{\{.+\}\})`).ReplaceAll(bs, []byte("$1"))
	bs = regexp.MustCompile(`\/\*\*\s*(\{\{.+\}\})\s*\*\*\/`).ReplaceAll(bs, []byte("$1"))

	t, err := template.New(*pathIn).Parse(string(bs))

	if err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	var b bytes.Buffer

	if err := t.Execute(b, structValue.Interface()); err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	if b, err = format.Source(b); err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(*pathOut, b.Bytes(), 0644); err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	osFile, err := os.OpenFile(*pathOut, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	osFile.Close()

	if err := osFile.Close(); err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	/*var fileSet = token.NewFileSet()
	var astFile, err = parser.ParseFile(fileSet, *pathIn, nil, parser.ParseComments)

	if err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	osFile, err := os.OpenFile(*pathOut, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	if err := macro.Generate(fileSet, astFile, map[string][]string(v)); err != nil {
		osFile.Close()
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	var c = &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}

	if err := c.Fprint(osFile, fileSet, astFile); err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}

	if err := osFile.Close(); err != nil {
		fmt.Println("goo:", err)
		os.Exit(1)
	}*/
}
