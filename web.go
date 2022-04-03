package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	http.HandleFunc("/", sayhello)
	http.HandleFunc("/bye", saybye)
	http.HandleFunc("/create", create)
	http.HandleFunc("/edit", edit)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/test/", HelloServer)
	http.ListenAndServe(":8000", nil)

}

type City struct {
	ID       string
	NameCity string
}

func create(w http.ResponseWriter, r *http.Request) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS city (id INTEGER PRIMARY KEY, name_city TEXT)")
	statement.Exec()

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		input := r.FormValue("city")
		statement, _ = database.Prepare("INSERT INTO city (name_city) VALUES (?)")
		statement.Exec(input)

	}
}

func edit(w http.ResponseWriter, r *http.Request) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS city (id INTEGER PRIMARY KEY, name_city TEXT)")
	statement.Exec()

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		var id, input string
		input = r.FormValue("input_data")
		id = r.FormValue("edit_data")
		fmt.Println(input)
		fmt.Println(id)
		statement, _ = database.Prepare(fmt.Sprintf(`UPDATE city SET name_city=%v WHERE ID = %v`, input, id))
		statement.Exec()

	}
}

func delete(w http.ResponseWriter, r *http.Request) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS city (id INTEGER PRIMARY KEY, name_city TEXT)")
	statement.Exec()

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		input := r.FormValue("del")
		statement, _ = database.Prepare(fmt.Sprintf(`DELETE FROM city WHERE  id  = %v`, input))
		statement.Exec()

	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS city (id INTEGER PRIMARY KEY, name_city TEXT)")
	statement.Exec()

	rows, _ := database.Query("SELECT id,name_city FROM city")

	var (
		cities []City
		city   City
	)

	for rows.Next() {
		rows.Scan(&city.ID, &city.NameCity)
		cities = append(cities, city)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", cities)

}

func saybye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Пока!")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	title := strings.Split(r.URL.Path[1:], "/")

	last_chars_url := title[len(title)-1]

	fmt.Fprintf(w, "<title>%v</title><h1>%v</h1>", last_chars_url, last_chars_url)
}
