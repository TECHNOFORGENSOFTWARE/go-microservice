package controllers

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
