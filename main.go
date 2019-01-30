package main

// jq . jsonfile, jq .abc jsonfile (jq .["abc"] jsonfile), jq .["ab.c"] jsonfile
// xq . xmlfile ...

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func main() {
	raw := flag.Bool("r", false, "raw output instead of xml")
	flag.Parse()
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if len(flag.Args()) < 1 || info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Usage echo '<xml>...</xml>' | xq OPERATION")
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)

	selector := flag.Args()[0]

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(reader); err != nil {
		panic(err)
	}

	values := make([]string, 0)
	subdoc := etree.NewDocument()
	for _, e := range doc.FindElements(selector) {
		subdoc.AddChild(e)
		if len(e.Text()) > 0 {
			values = append(values, e.Text())
		}
	}

	if *raw {
		for _, e := range values {
			fmt.Println(e)
		}
	} else {
		subdoc.Indent(2)
		subdoc.WriteTo(os.Stdout)
	}
}
