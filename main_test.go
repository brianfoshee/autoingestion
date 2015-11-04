package main

import "testing"

func TestProcessArgs(t *testing.T) {
	tests := struct {
		Args []string
		Params
	}{}

	for _, t := range tests {
		p := processArgs(t.Args)
		if p.Params != t.Params {
			t.Errorf("expected (%+v) actual (%+v)", t.Params, p.Params)
		}
	}
}
