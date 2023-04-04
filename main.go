package main

import (
	"go-student/config"
	"go-student/controllers/homecontroller"
	"log"
	"net/http"
)



func main() {
	config.ConnectDB()

	// home 
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server Running on port 3000")
	http.ListenAndServe(":3000", nil)
}

