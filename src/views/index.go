package views

import (
	"data"
	"html/template"
	"net/http"
)

/*
Index is the main view of the app
Displays the unique visitors
*/
func Index(w http.ResponseWriter, r *http.Request) {
	templatedata := data.FetchData(w, r)
	t := template.New("index")
	t, _ = t.ParseFiles("./templates/test.html")
	t.ExecuteTemplate(w, "index", templatedata)
}
