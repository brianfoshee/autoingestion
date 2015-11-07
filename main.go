package main

import (
	"errors"
	"fmt"
	"log"
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

	Properties properties.Properties
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

	p, err := processArgs(args)
	if err != nil {
		fmt.Println("error processing args", err)
		return
	}
	log.Printf("got args %+v", p)

}

func processArgs(args []string) (Params, error) {
	p := Params{}
	if len(args) < 5 {
		return p, errors.New("not enough args")
	}

	p.PropertiesFilePath = args[0]
	p.Properties = properties.NewPropertiesFromFile(p.PropertiesFilePath)
	p.VendorID = args[1]
	return p, nil
}
