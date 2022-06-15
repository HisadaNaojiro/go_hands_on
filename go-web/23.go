package main

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"text/template"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("go-web/templates/"+fname+".html", "go-web/templates/head.html", "go-web/templates/foot.html")
	return tmps
}

func index(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "index",
		Title:    "Index",
		Message:  "This is Top page.",
	}
	er := page("index2").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func hello(w http.ResponseWriter, rq *http.Request) {
	data := []string{"One", "Two", "Three"}

	item := struct {
		Title string
		Data  []string
	}{
		Title: "Hello",
		Data:  data,
	}
	er := page("hello7").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq)
	})
	http.ListenAndServe("", nil)
}
