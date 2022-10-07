package main

import (
	"net/http"

	"github.com/TECHNOFORGENSOFTWARE/go-microservice/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
