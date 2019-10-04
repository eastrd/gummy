package main

import ()

type G struct {
	Method string `yaml:"method"`
	Route  string `yaml:"route"`
	Resp   struct {
		Status  int               `yaml:"status"`
		Body    string            `yaml:"body"`
		Headers map[string]string `yaml:"headers"`
	} `yaml:"response"`
}

func (g G) populateDefault() G {
	// Some fields are filled with default values if left empty
	if g.Resp.Status == 0 {
		// Status default to be 200
		g.Resp.Status = 200
	}
	if g.Method == "" {
		// Method default to be GET
		g.Method = "GET"
	}
	return g
}

func (g G) validateRules() bool {
	// Validate each field
	return true
}
