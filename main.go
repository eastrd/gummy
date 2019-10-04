package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("sample.yaml")
	if err != nil {
		panic(err)
	}

	// Parse the yaml
	var gs []G
	err = yaml.Unmarshal(b, &gs)
	if err != nil {
		panic(err)
	}

	for idx, g := range gs {
		gs[idx] = g.populateDefault()
		if !gs[idx].validateRules() {
			panic("error validating rules")
		}
	}

	fmt.Printf("%+v\n", gs)

	// Generate Golang code
	code := ""
	code += HEAD + "\n"
	for _, g := range gs {
		replHandler := strings.ReplaceAll(HANDLER, "{{STATUS}}", strconv.Itoa(g.Resp.Status))
		replHandler = strings.ReplaceAll(replHandler, "{{ROUTE}}", g.Route)
		replHandler = strings.ReplaceAll(replHandler, "{{BODY}}", g.Resp.Body)
		if len(g.Resp.Headers) > 0 {
			headers := ""
			for k, v := range g.Resp.Headers {
				newHeader := strings.ReplaceAll(HEADER, "{{KEY}}", k)
				newHeader = strings.ReplaceAll(newHeader, "{{VALUE}}", v)
				headers += newHeader + "\n\t\t"
			}
			// Insert header lines to existing code
			replHandler = strings.ReplaceAll(replHandler, "{{HEADERS}}", headers)

		} else {
			// Remove the {{HEADERS}} string as there are no headers
			replHandler = strings.ReplaceAll(replHandler, "{{HEADERS}}", "")
		}

		code += replHandler + "\n"
	}

	code += TAIL

	fmt.Println("Generated Code:\n" + code)
	err = ioutil.WriteFile("generated/gen.go", []byte(code), 0644)
	if err != nil {
		panic(err)
	}
}
