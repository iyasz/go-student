package main

import (
	"go-student/config"
	"go-student/controllers/homecontroller"
	"go-student/controllers/studentcontroller"
	"log"
	"net/http"
)



func main() {
	config.ConnectDB()

	// home 
	http.HandleFunc("/", homecontroller.Welcome)

	// student 
	http.HandleFunc("/student", studentcontroller.Index)
	http.HandleFunc("/student/create", studentcontroller.Create)
	http.HandleFunc("/student/edit", studentcontroller.Edit)
	
	// http.HandleFunc("/student/edit", studentcontroller.Edit)
	// http.HandleFunc("/student/", studentcontroller.Index)

	log.Println("Server Running on port 3000")
	http.ListenAndServe(":3000", nil)
}

