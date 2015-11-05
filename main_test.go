package main

import (
	"testing"

	"github.com/brianfoshee/autoingestion/properties"
)

func TestProcessArgs(t *testing.T) {
	tests := []struct {
		Args       []string
		Valid      bool
		Properties properties.Properties
	}{
		{
			Args:       []string{"properties"},
			Valid:      false,
			Properties: properties.Properties{},
		},
	}

	for _, c := range tests {
		p, err := processArgs(c.Args)
		if err != nil && c.Valid {
			t.Errorf("expected an error")
		}
		if p.Properties != c.Properties {
			t.Errorf("expected (%+v) actual (%+v)", c.Properties, p.Properties)
		}
	}
}
