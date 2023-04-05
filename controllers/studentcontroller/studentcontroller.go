package studentcontroller

import (
	"go-student/models/studentmodel"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	students := studentmodel.GetAll()

	data := map[string]any {
		"student": students,
	}
	
	temp, err := template.ParseFiles("views/student/index.html")

	if err != nil{
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/student/create.html")

	if err != nil{
		panic(err)
	}

	temp.Execute(w, nil)
}

