package main

import (
	"log"
	"net/http"
	"os"
	"views"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("style"))))
	http.HandleFunc("/", views.Index)
	log.Println("Server started")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
