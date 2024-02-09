package main

import (
	"flag"
	"fmt"

	"github.com/jpbruinsslot/resume"
)

const (
	VERSION = "1.1.0"
	URL     = "github.com/jpbruinsslot/resume"

	USAGE = `NAME:
    resume - resume server that can be accessed by a curl command.

USAGE:
    resume [options]

VERSION:
    %s

WEBSITE:
    %s

OPTIONS:
    -port       port to serve on (default: 80)
    -path       optional url path (default: /)
    -template   filename of template (default: resume.tmpl)
    -user       username used for authentication (default: "")
    -pass       password used for authentication (default: "")
`
)

func main() {
	flgPort := flag.String(
		"port", "80", "port to serve on",
	)

	flgPath := flag.String(
		"path", "/", "optional url path",
	)

	flgTemplate := flag.String(
		"template", "resume.tmpl", "filename of template",
	)

	flgUser := flag.String(
		"user", "", "username used for authentication",
	)

	flgPass := flag.String(
		"pass", "", "password used for authentication",
	)

	flag.Usage = func() {
		fmt.Printf(USAGE, VERSION, URL)
	}

	flag.Parse()

	resume.InitServer(*flgPort, *flgPath, *flgTemplate, *flgUser, *flgPass)
}
