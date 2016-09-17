package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/willfaught/goo"
)

func handle(err error, s string) {
	if err == nil {
		return
	} else if s == "" {
		fmt.Println("goo: error:", err)
	} else {
		fmt.Printf("goo: error: %v: %v", s, err)
	}

	os.Exit(1)
}

func main() {
	var flagHelp = flag.Bool("help", false, "Print help.")
	var flagIn = flag.String("in", "", "Input file path. Required.")
	var flagJSON = flag.String("json", "null", "JSON data. Conflicts with field flag.")
	var flagOut = flag.String("out", "", "Output file path. Required.")

	flag.Parse()

	if *flagHelp {
		flag.PrintDefaults()

		return
	}

	if *flagIn == "" {
		handle(fmt.Errorf("in flag is required"), "")
	}

	if *flagOut == "" {
		handle(fmt.Errorf("out flag is required"), "")
	}

	var j interface{}

	handle(json.Unmarshal([]byte(*flagJSON), &j), "invalid json")

	var bs, err = ioutil.ReadFile(*flagIn)

	handle(err, fmt.Sprintf("cannot read %v", *flagIn))
	bs, err = goo.Macro(*flagIn, bs, j)
	handle(err, "macro")
	handle(ioutil.WriteFile(*flagOut, bs, 0644), fmt.Sprintf("cannot write %v", *flagOut))
}
