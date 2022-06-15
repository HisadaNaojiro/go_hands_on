package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	html := `<html><body>
<h1>Hello</h1>
<p> This is GO-server!!</p>
</body></html>`
	tf, er := template.New("index").Parse(html)
	if er != nil {
		log.Fatal(er)
	}
	hh := func(w http.ResponseWriter, rq *http.Request) {
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}
	http.HandleFunc("/hello", hh)
	http.ListenAndServe("", nil)
}
