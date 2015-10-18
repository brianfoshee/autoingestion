package properties

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Properties struct {
	Email    string
	Password string
}

func NewPropertiesFromFile(name string) Properties {
	p := Properties{}
	p.fromFile(name)
	return p
}

func (p *Properties) fromFile(name string) {
	str := p.readFile(name)
	if str == "" {
		return
	}

	p.extractEmail(str)

	p.extractPassword(str)
}

func (p *Properties) extractEmail(str string) {
	// find the account email address
	re := regexp.MustCompile("(?m:^userID = )((?m)\\w+@\\w+\\.com$)")
	m := re.FindStringSubmatch(str)
	// first match is the whole regex. Second is the subgroup match.
	if len(m) != 2 {
		fmt.Println("did not find a match", m)
		return
	}
	p.Email = m[1]
}

func (p *Properties) extractPassword(str string) {
	// find the account password
	re := regexp.MustCompile("(?m:^password = )((?m).+$)")
	m := re.FindStringSubmatch(str)
	// first match is the whole regex. Second is the subgroup match.
	if len(m) != 2 {
		fmt.Println("did not find a match", m)
		return
	}
	p.Password = m[1]
}

func (p *Properties) readFile(fn string) string {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("error opening properties file", err)
		return ""
	}
	defer file.Close()

	// Get the size of the file
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("could not read file stats", err)
		return ""
	}
	nb := fi.Size()

	// Read the file into memory
	b := make([]byte, nb)
	rb, err := file.Read(b)
	if err != nil {
		fmt.Println("could not read file", err)
		return ""
	}
	if int64(rb) != nb {
		fmt.Println("did not read in as many bytes as size of file", err)
		return ""
	}

	return strings.TrimSpace(string(b))
}
