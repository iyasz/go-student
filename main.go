package main

import (
	"go-student/config"
	"log"
	"net/http"
)



func main() {
	config.ConnectDB()

	log.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}

