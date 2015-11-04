package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brianfoshee/autoingestion/properties"
)

/*
** CLI format **
java Autoingestion <properties_filename> <vendor_id> <report_type> <date_type> <report_subtype> <date>

** Args details **
Parameter | Values | Notes
properties_filename | Name of the properties | file Make sure the file extension is “.properties”. For example, jane_doe.properties.
vendor_id | Your unique vendor number | The vendor ID for which you want to download the report. For example, 80012345.
report_type | Sales or Newsstand |
date_type | Daily, Weekly, Monthly, Yearly
report_subtype | Summary, Detailed, or Opt-In | Opt-In only applies to Sales report.
date (optional) | YYYYMMDD (Daily or Weekly)  | YYYYMM (Monthly) YYYY (Yearly) | The date of the report you are requesting. Date parameter is optional. If it is not provided, you will get the latest report available.
*/

type Params struct {
	PropertiesFilePath string
	VendorID           string
	ReportType         string
	DateType           string
	ReportSubType      string
	Date               time.Time

	Properties
}

func main() {
	args := os.Args

	for _, x := range args {
		fmt.Println(x)
	}
	/*
	   - read in args
	   - make sure there are enough
	   - store them in a struct
	   - read in the properties file to get username/password
	   - modify the runAutoingestion function to take that struct as variable
	   - use those values in the form data
	*/

	if len(args) < 5 {
		fmt.Println("not enough args")
		return
	}

	p, err := processArgs(args)
	if err != nil {
		fmt.Println("error processing args", err)
		return
	}

}

func processArgs(args []string) Params {
	p := Params{}
	p.Properties = properties.NewPropertiesFromFile(p.PropertiesFilePath)
}
