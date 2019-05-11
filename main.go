package main

import (
	"flag"

	"github.com/erroneousboat/resume/pkg/resume"
)

func main() {
	flgPort := flag.String("port", "80", "port to serve on")
	flgTemplate := flag.String("template", "resume.tmpl", "filename of template")
	flgUser := flag.String("user", "", "username used for authentication")
	flgPass := flag.String("pass", "", "password used for authentication")
	flag.Parse()

	resume.InitServer(*flgPort, *flgTemplate, *flgUser, *flgPass)
}
