package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type MyHandler struct {
}

func (j *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello</h1>")
}
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "heelo")
}

func main() {
	http.HandleFunc("/top", top)
	//log.Fatal(http.ListenAndServe(":8080", &MyHandler{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
