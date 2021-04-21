package main

import (
	"log"
	"net/http"

	"gawds/src/middleware"
	"gawds/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Server Running"))
	})
	r.Use(middleware.Middleware)
	r.HandleFunc("/login", routes.Login).Methods("POST")
	r.HandleFunc("/logout", routes.).Methods("POST")
	r.HandleFunc("/register", routes.CreateUser).Methods("POST")
	r.HandleFunc("/profile", routes.GetUser).Methods("GET")

	log.Println("Server Stated")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Error Occured", err)
	}
}
