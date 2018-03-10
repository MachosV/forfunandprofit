package templater

import (
	"html/template"
	"io/ioutil"
	"log"
	"strings"
)

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	files, err := ioutil.ReadDir("./templates/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") == false {
			continue
		}
		t := template.New(filename)
		t, err = t.ParseFiles("./templates/"+filename, "./templates/base.html")
		log.Println(t)
		templates[filename] = t
	}
}

func GetTemplate(templateName string) *template.Template {
	return templates[templateName]
}
