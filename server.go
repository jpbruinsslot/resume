package resume

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// InitServer starts the resume server
func InitServer(flgPort, flgPath, flgTemplate, flgUser, flgPass string) {
	colors := GetANSIColors()

	tmpl, err := template.ParseFiles(flgTemplate)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc(flgPath, func(w http.ResponseWriter, r *http.Request) {
		// authenticate when user and pass flags are set
		if flgUser != "" && flgPass != "" {
			if auth := Authenticated(r, flgUser, flgPass); auth == false {
				w.Header().Set("WWW-Authenticate", "password protected")
				w.WriteHeader(401)
				w.Write([]byte("401 Unauthorized\n"))
				return
			}
		}

		// create buffer to write results to
		buf := new(bytes.Buffer)

		// execute template with colors context
		err = tmpl.Execute(buf, colors)
		if err != nil {
			log.Fatal(err)
		}

		// output results
		fmt.Fprint(w, buf.String())
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", flgPort), nil))
}
