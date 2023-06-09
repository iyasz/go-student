package studentcontroller

import (
	"go-student/entities"
	"go-student/models/studentmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {

	students := studentmodel.GetAll()

	data := map[string]any{
		"student": students,
	}

	temp, err := template.ParseFiles("views/student/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request){
	temp, err := template.ParseFiles("views/student/detail.html")

	if err != nil {
		panic(err)
	}

	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	student := studentmodel.Detail(id)
	data := map[string]any{
		"student": student,
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		temp, err := template.ParseFiles("views/student/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var student entities.Student

		student.Nama = r.FormValue("nama")
		student.Nis = r.FormValue("nis")
		student.Umur = r.FormValue("umur")
		student.Gender = r.FormValue("gender")
		student.Telp = r.FormValue("telp")
		student.Alamat = r.FormValue("alamat")
		student.CreatedAt = time.Now()

		if ok := studentmodel.Store(student); !ok {
			temp, _ := template.ParseFiles("views/student/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/student", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request){
	if r. Method == "GET" {
		temp, err := template.ParseFiles("views/student/edit.html")

		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		student := studentmodel.Detail(id)
		data := map[string]any{
			"student": student,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var student entities.Student

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		student.Nama = r.FormValue("nama")
		student.Nis = r.FormValue("nis")
		student.Umur = r.FormValue("umur")
		student.Gender = r.FormValue("gender")
		student.Telp = r.FormValue("telp")
		student.Alamat = r.FormValue("alamat")

		if ok := studentmodel.Update(id, student); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/student", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := studentmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/student", http.StatusSeeOther)
}
