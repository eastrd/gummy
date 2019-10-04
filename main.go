package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
		code += strings.ReplaceAll(strings.ReplaceAll(HANDLER, "{{ROUTE}}", g.Route), "{{BODY}}", g.Resp.Body) + "\n"
	}

	code += TAIL

	fmt.Println("Generated Code:\n" + code)
	err = ioutil.WriteFile("generated/gen.go", []byte(code), 0644)
	if err != nil {
		panic(err)
	}
}
