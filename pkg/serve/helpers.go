package serve

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, name, tpl string, values interface{}) {
	fMap := template.FuncMap{
		"mod": func(a, shift, mod int) bool { return (a+shift)%mod == 0 },
	}

	t, err := template.New(name).Funcs(fMap).Parse(tpl)
	if err != nil {
		log.Printf("error parsing %s template: %v", name, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.New("header").Parse(baseTpl)
	t.Execute(w, values)
}
