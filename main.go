package main

// jq . jsonfile, jq .abc jsonfile (jq .["abc"] jsonfile), jq .["ab.c"] jsonfile
// xq . xmlfile ...

import (
	"flag"
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func main() {
	file := flag.String("i", "config.xml", "input file")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Usage:\txq [-i FILE] SELECTOR\nor:\techo '<xml>...</xml>' | xq SELECTOR")
		os.Exit(1)
	}
	selector := flag.Args()[0]

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(*file); err != nil {
		panic(err)
	}

	subdoc := etree.NewDocument()
	for _, e := range doc.FindElements(selector) {
		//fmt.Printf("%s: %s\n", e.Tag, e.Text())
		subdoc.AddChild(e)
	}

	subdoc.Indent(2)
	subdoc.WriteTo(os.Stdout)
}
