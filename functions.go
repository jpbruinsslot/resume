package resume

import (
	"strings"
	"text/template"
)

// FuncMap is the template.FuncMap for the templates
var Funcs = template.FuncMap{

	// "wrap" is a function that wraps a string to a given width
	"wrap": func(s string, width int) string {
		words := strings.Fields(strings.TrimSpace(s))
		if len(words) == 0 {
			return s
		}

		wrapped := words[0]
		spaceLeft := width - len(wrapped)
		for _, word := range words[1:] {
			if len(word)+1 > spaceLeft {
				// remove space when padding not needed
				wrapped += "\n" + " " + word
				spaceLeft = width - len(word)
			} else {
				wrapped += " " + word
				spaceLeft -= 1 + len(word)
			}
		}

		return wrapped
	},
}
