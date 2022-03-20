package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	http.HandleFunc("/", sayhello)
	http.HandleFunc("/bye", saybye)
	http.HandleFunc("/test/", HelloServer)
	http.ListenAndServe(":8000", nil)

}

func sayhello(w http.ResponseWriter, r *http.Request) {

	lst := []string{}

	for i := 0; i < 100; i++ {

		lst = append(lst, strconv.Itoa(i))

	}

	tpl.ExecuteTemplate(w, "index.gohtml", lst)

}

func saybye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Пока!")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	title := strings.Split(r.URL.Path[1:], "/")

	last_chars_url := title[len(title)-1]

	fmt.Fprintf(w, "<title>%v</title><h1>%v</h1>", last_chars_url, last_chars_url)
}
