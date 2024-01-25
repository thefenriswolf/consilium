package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed resources/*
var resources embed.FS

var tpl *template.Template
var name = "john"

const PORT string = ":6477"
const ADDRESS string = "http://localhost" + PORT

func handleInteractiveTestPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loading interactive test page")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loading main page")
	fmt.Fprintf(w, "Welcome to the main page!")
}

func handleTestPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loading static test page")
	fmt.Fprintf(w, "found the test page")
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("templated loaded")
	err := tpl.ExecuteTemplate(w, "welcome.html", name)
	if err != nil {
		fmt.Println(err)
	}
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	buf, err := fs.ReadFile(resources, "resources/assets/dragon.png")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
}

func main() {
	tpl, _ = tpl.ParseFS(resources, "resources/templates/welcome.html")
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/test", handleTestPage)
	http.HandleFunc("/itest", handleInteractiveTestPage)
	http.HandleFunc("/temp", templateHandler)
	http.HandleFunc("/dragon", handleImage)

	fmt.Printf("Listening on: http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
