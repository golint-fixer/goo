package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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

func handle(err error, s string) {
	if err == nil {
		return
	} else if s == "" {
		fmt.Println("goo: error:", err)
	} else {
		fmt.Printf("goo: error: %v: %v\n", s, err)
	}

	os.Exit(1)
}

func main() {
	var flagIn = flag.String("in", "", "Input file path. Defaults to standard in.")
	var flagJSON = flag.String("json", "null", "Macro data as JSON.")
	var flagOut = flag.String("out", "", "Output file path. Defaults to standard out.")
	var flagPreprocess = flag.Bool("preprocess", false, "Do not process and format the macro.")
	var flagProcess = flag.Bool("process", false, "Do not format the macro.")

	flag.Parse()

	var j interface{}

	handle(json.Unmarshal([]byte(*flagJSON), &j), "invalid json")
	convert(j)

	var fileIn, fileOut *os.File
	var nameIn, nameOut string
	var err error

	if *flagIn == "" {
		nameIn = "standard in"
		fileIn = os.Stdin
	} else {
		nameIn = *flagIn
		fileIn, err = os.Open(*flagIn)
		handle(err, fmt.Sprintf("cannot open %v", *flagIn))
	}

	if *flagOut == "" {
		nameOut = "standard out"
		fileOut = os.Stdout
	} else {
		nameOut = *flagOut
		fileOut, err = os.OpenFile(*flagOut, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		handle(err, fmt.Sprintf("cannot open %v", *flagOut))
	}

	var bs []byte

	bs, err = ioutil.ReadAll(fileIn)
	handle(err, fmt.Sprint("cannot read", nameIn))
	bs = goo.MacroPreprocess(bs)

	if *flagPreprocess {
		_, err = fileOut.Write(bs)
		handle(err, fmt.Sprint("cannot write", nameOut))

		return
	}

	bs, err = goo.MacroProcess(*flagIn, bs, j)
	handle(err, "cannot process macro")

	if *flagProcess {
		_, err = fileOut.Write(bs)
		handle(err, fmt.Sprint("cannot write", nameOut))

		return
	}

	bs, err = goo.MacroFormat(bs)
	handle(err, "cannot format macro")

	_, err = fileOut.Write(bs)
	handle(err, fmt.Sprint("cannot write", nameOut))
}
