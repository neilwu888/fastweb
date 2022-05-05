package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"web/password"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloHandler is invoked")
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
	fmt.Println("helloHandler served")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func layoutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/layout.html", "static/header.html", "static/footer.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func toolsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/tools/main.html", "static/header.html", "static/footer.html"))
	tmpl.Execute(w, nil)
}

// func passwordHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Loading password list")
// 	tmpl := template.Must(template.ParseFiles("views/tools/password/main.html", "static/header.html", "static/footer.html"))
// 	tmpl.Execute(w, nil)
// }

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	http.HandleFunc("/hello", helloHandler)             // Update this line of code
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/layout", layoutHandler)
	http.HandleFunc("/tools", toolsHandler)
	// http.HandleFunc("/tools/password", password.PasswordHandler)
	password.Route()

	fmt.Printf("Starting server at port 9090\n")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
