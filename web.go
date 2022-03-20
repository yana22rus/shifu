package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
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

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS city(id INTEGER PRIMARY KEY, name_city TEXT)")
	statement.Exec()

	rows, _ := database.Query("SELECT name_city FROM city")

	var city string

	for rows.Next() {
		rows.Scan(&city)
		lst = append(lst, city)
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

