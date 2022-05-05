package password

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"web/model"
	"web/utils"
)

//
func Route() {
	http.HandleFunc("/tools/password", mainHandler)
	http.HandleFunc("/tools/password/create", createForm)
	http.HandleFunc("/tools/password/save", saveForm)
	http.HandleFunc("/tools/password/edit", editForm)
	http.HandleFunc("/tools/password/delete", deleteForm)
	http.HandleFunc("/tools/password/update", updateForm)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Loading password list")
	tmpl := template.Must(template.ParseFiles("views/tools/password/main.html", "static/header.html", "static/footer.html"))
	lisOfPassword := utils.ReadPassword()
	fmt.Println(lisOfPassword)
	tmpl.Execute(w, lisOfPassword)
}

func createForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Loading password list")
	tmpl := template.Must(template.ParseFiles("views/tools/password/createForm.html", "static/header.html", "static/footer.html"))
	tmpl.Execute(w, nil)
}

func editForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Editing password list")
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	fmt.Println("Editing password,id=" + string(id))
	list := utils.ReadPassword()
	fmt.Println(list)
	var targetPassword model.TsPassword
	for _, p := range list.List {
		if id == p.Id {
			targetPassword = p
		}
	}

	fmt.Println(targetPassword)
	tmpl := template.Must(template.ParseFiles("views/tools/password/editForm.html", "static/header.html", "static/footer.html"))
	error := tmpl.Execute(w, targetPassword)
	fmt.Println(error)
}

func saveForm(w http.ResponseWriter, r *http.Request) {
	var data model.TsPassword
	data.Id = -1
	data.Name = r.FormValue("name")
	data.Password = r.FormValue("password")
	data.Details = r.FormValue("details")
	utils.SavePassword(data)
	mainHandler(w, r)
}

func updateForm(w http.ResponseWriter, r *http.Request) {
	var data model.TsPassword
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	data.Id = id
	data.Name = r.FormValue("name")
	data.Password = r.FormValue("password")
	data.Details = r.FormValue("details")
	utils.SavePassword(data)
	mainHandler(w, r)
}

func deleteForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	fmt.Println("Deleting password,name=" + string(id))
	utils.DeletePassword(id)
	mainHandler(w, r)
}
