package main

// jq . jsonfile, jq .abc jsonfile (jq .["abc"] jsonfile), jq .["ab.c"] jsonfile
// xq . xmlfile ...

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func xmlToMap(r io.Reader) map[string]string {
	// result
	m := make(map[string]string)
	// the current value stack
	values := make([]string, 0)
	// parser
	p := xml.NewDecoder(r)
	for token, err := p.Token(); err == nil; token, err = p.Token() {
		switch t := token.(type) {
		case xml.CharData:
			// push
			values = append(values, string([]byte(t)))
		case xml.EndElement:
			if t.Name.Local == "langs" {
				continue
			}
			m[t.Name.Local] = values[len(values)-1]
			// pop
			values = values[:len(values)]
		}
	}
	// done
	return m
}

func parse(r io.Reader) map[string]string {
	decoder := xml.NewDecoder(r)
	//solutions := make([]string, 0, 0)
	solutions := make(map[string]string)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "Paragraph" {
				// Get the next token after the Paragraph start element, which will be the tag contents
				//innerText, ok := decoder.Token().(xml.CharData)
				innerText, err := decoder.Token()
				//if !ok {
				if err != nil {
					continue
				}
				//solutions = append(solutions, string(innerText))
				//solutions = append(solutions, string(innerText.(xml.CharData)))
				solutions["x"] = string(innerText.(xml.CharData))
			}
		}
	}
	return solutions
}

func main() {
	file, err := os.Open("config.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//xmlReader := bytes.NewReader([]byte(your_xml_as_a_string_here))
	//yourPinnacleLineFeed := new(Pinnacle_Line_Feed)
	//if err := xml.NewDecoder(xmlReader).Decode(yourPinnacleLineFeed); err != nil {
	//return // or log.Panic(err.Error()) if in main
	//}
	//r := strings.NewReader(XML)
	//m := xmlToMap(file)
	m := parse(file)
	fmt.Println(m)
}
