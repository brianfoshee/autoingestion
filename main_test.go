package main

import (
	"testing"

	"github.com/brianfoshee/autoingestion/properties"
)

func TestProcessArgs(t *testing.T) {
	tests := []struct {
		Args   []string
		Valid  bool
		Params Params
	}{
		{
			Args:   []string{"properties"},
			Valid:  false,
			Params: Params{},
		},
		{
			Args:  []string{"autoingestion.properties.example", "123", "Sales", "Daily", "Summary"},
			Valid: true,
			Params: Params{
				PropertiesFilePath: "autoingestion.properties",
				VendorID:           "123",
				ReportType:         "Sales",
				DateType:           "Daily",
				ReportSubType:      "Summary",
				Properties: properties.Properties{
					Email:    "user@example.com",
					Password: "P@$$w0rd",
				},
			},
		},
	}

	for _, c := range tests {
		p, err := processArgs(c.Args)
		if err != nil && c.Valid {
			t.Errorf("expected an error")
		}
		if p != c.Params {
			t.Errorf("expected (%+v) actual (%+v)", c.Params, p)
		}
	}
}
